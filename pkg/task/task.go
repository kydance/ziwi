package task

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
)

type Result struct {
	Data any
	Err  error
}

type ProcessorStats struct {
	completedTasks uint64
	errCnt         uint64
	activeWorkers  int32
}

type Processor struct {
	maxWorkers int
	workerPool chan struct{}

	errHandler func(error)
	retHandler func(any)

	stats ProcessorStats
}

type ProcessorOption func(*Processor)

func NewTaskProcessor(options ...ProcessorOption) *Processor {
	// 默认使用 CPU 核数作为工作协程数量
	p := &Processor{
		maxWorkers: runtime.GOMAXPROCS(0),
	}

	// 应用配置选项
	for _, option := range options {
		option(p)
	}

	// Init worker pool
	p.workerPool = make(chan struct{}, p.maxWorkers)
	for i := 0; i < p.maxWorkers; i++ {
		p.workerPool <- struct{}{}
	}

	return p
}

// WithMaxWorkers 设置最大工作协程数量
func WithMaxWorkers(n int) ProcessorOption {
	return func(p *Processor) {
		p.maxWorkers = n
	}
}

// WithErrorHandler 设置错误处理函数
func WithErrorHandler(handler func(error)) ProcessorOption {
	return func(p *Processor) {
		p.errHandler = handler
	}
}

// WithRetHandler 设置结果处理函数
func WithRetHandler(handler func(any)) ProcessorOption {
	return func(p *Processor) {
		p.retHandler = handler
	}
}

func (p *Processor) ProcessChunks(ctx context.Context, data []any,
	processFunc func([]any) (any, error),
) error {
	if len(data) == 0 {
		return nil
	}

	// 计算每个工作协程的数据量
	chunkSize := (len(data) + p.maxWorkers - 1) / p.maxWorkers // 向上取整
	chunks := make([][]any, 0, p.maxWorkers)

	// allocate chunks
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}

	// 创建错误通道
	errChan := make(chan error, len(chunks))
	resultChan := make(chan Result, len(chunks))

	var wg sync.WaitGroup
	wg.Add(len(chunks))

	// Handle each chunk
	for i := range chunks {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-p.workerPool:
			go func(chunk []any) {
				defer func() {
					p.workerPool <- struct{}{} // 释放工作协程槽位
					wg.Done()
				}()

				// Update active workers
				atomic.AddInt32(&p.stats.activeWorkers, 1)
				defer atomic.AddInt32(&p.stats.activeWorkers, -1)

				result, err := processFunc(chunk)
				if err != nil {
					atomic.AddUint64(&p.stats.errCnt, 1)
					if p.errHandler != nil {
						p.errHandler(err)
					}

					errChan <- err
					return
				}

				atomic.AddUint64(&p.stats.completedTasks, 1)
				if p.retHandler != nil {
					p.retHandler(result)
				}
				resultChan <- Result{
					Data: result,
					Err:  nil,
				}
			}(chunks[i])
		}
	}

	// 等待所有工作协程完成
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
		close(errChan)
		close(resultChan)
	}()

	// 等待完成或上下文被取消
	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-done:
		return nil
	}
}

// Stats returns 统计信息
func (p *Processor) Stats() ProcessorStats {
	return ProcessorStats{
		activeWorkers:  atomic.LoadInt32(&p.stats.activeWorkers),
		completedTasks: atomic.LoadUint64(&p.stats.completedTasks),
		errCnt:         atomic.LoadUint64(&p.stats.errCnt),
	}
}
