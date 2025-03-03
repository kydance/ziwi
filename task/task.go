package task

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
)

// TaskResult holds the result of a task along with any associated error
type TaskResult struct {
	Result any
	Error  error
}

// ProcessorMetrics keeps track of key metrics for the task processor
type ProcessorMetrics struct {
	completedTasks uint64 // number of successfully completed tasks
	errorCount     uint64 // number of errors encountered
	activeWorkers  int32  // current active worker count
}

// TaskProcessor is responsible for managing task execution with a pool of workers
type TaskProcessor struct {
	maxWorkers int           // maximum number of concurrent workers
	workerPool chan struct{} // worker pool semaphore to limit concurrency

	errorHandler  func(error) // function to handle task errors
	resultHandler func(any)   // function to handle task results

	metrics ProcessorMetrics // processor's performance metrics
}

// ProcessorOption defines the functional option type for TaskProcessor configuration
type ProcessorOption func(*TaskProcessor)

// NewTaskProcessor creates a new TaskProcessor with the provided options
func NewTaskProcessor(options ...ProcessorOption) *TaskProcessor {
	// Default worker count is the number of logical CPUs
	processor := &TaskProcessor{
		maxWorkers: runtime.GOMAXPROCS(0),
	}

	// Apply configuration options
	for _, option := range options {
		option(processor)
	}

	// Initialize worker pool
	processor.workerPool = make(chan struct{}, processor.maxWorkers)
	for i := 0; i < processor.maxWorkers; i++ {
		processor.workerPool <- struct{}{} // pre-fill the worker pool
	}

	return processor
}

// WithMaxWorkerCount configures the maximum number of workers for the TaskProcessor
func WithMaxWorkerCount(n int) ProcessorOption {
	return func(p *TaskProcessor) { p.maxWorkers = n }
}

// WithErrorHandler configures the error handler for tasks that encounter errors
func WithErrorHandler(handler func(error)) ProcessorOption {
	return func(p *TaskProcessor) { p.errorHandler = handler }
}

// WithResultHandler configures the result handler for tasks that return results
func WithResultHandler(handler func(any)) ProcessorOption {
	return func(p *TaskProcessor) { p.resultHandler = handler }
}

// ProcessInChunks processes tasks in chunks, distributing them among workers
func (p *TaskProcessor) ProcessInChunks(ctx context.Context, tasks []any, taskFunc func([]any) (any, error)) error {
	if len(tasks) == 0 {
		return nil // no tasks to process
	}

	// Determine chunk size based on the number of workers
	chunkSize := (len(tasks) + p.maxWorkers - 1) / p.maxWorkers // ceiling division
	taskChunks := make([][]any, 0, p.maxWorkers)

	// Split tasks into chunks
	for i := 0; i < len(tasks); i += chunkSize {
		end := i + chunkSize
		if end > len(tasks) {
			end = len(tasks)
		}
		taskChunks = append(taskChunks, tasks[i:end])
	}

	// Error and result channels for handling task outcomes
	errorChan := make(chan error, len(taskChunks))
	resultChan := make(chan TaskResult, len(taskChunks))

	var wg sync.WaitGroup
	wg.Add(len(taskChunks)) // each chunk is processed by one worker

	// Process each task chunk with a worker
	for i := range taskChunks {
		select {
		case <-ctx.Done():
			return ctx.Err() // context canceled, exit early

		case <-p.workerPool:
			// Launch a goroutine to handle this chunk
			go func(chunk []any) {
				defer func() {
					p.workerPool <- struct{}{} // release worker slot back to the pool
					wg.Done()                  // mark this worker as done
				}()

				// Track active workers
				atomic.AddInt32(&p.metrics.activeWorkers, 1)
				defer atomic.AddInt32(&p.metrics.activeWorkers, -1)

				// Execute the task function for this chunk
				result, err := taskFunc(chunk)
				if err != nil {
					// Track error count and invoke error handler
					atomic.AddUint64(&p.metrics.errorCount, 1)
					if p.errorHandler != nil {
						p.errorHandler(err)
					}
					errorChan <- err
					return
				}

				// Track completed tasks and invoke result handler
				atomic.AddUint64(&p.metrics.completedTasks, 1)
				if p.resultHandler != nil {
					p.resultHandler(result)
				}
				resultChan <- TaskResult{
					Result: result,
					Error:  nil,
				}
			}(taskChunks[i])
		}
	}

	// Wait for all workers to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()         // wait for all goroutines to finish
		close(done)       // signal completion
		close(errorChan)  // close error channel
		close(resultChan) // close result channel
	}()

	// Wait for either all tasks to complete or context cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-done:
		return nil
	}
}

// Metrics returns the current processor's performance metrics
func (p *TaskProcessor) Metrics() ProcessorMetrics {
	return ProcessorMetrics{
		activeWorkers:  atomic.LoadInt32(&p.metrics.activeWorkers),
		completedTasks: atomic.LoadUint64(&p.metrics.completedTasks),
		errorCount:     atomic.LoadUint64(&p.metrics.errorCount),
	}
}
