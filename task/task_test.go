package task

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/kydance/ziwi/log"
)

func TestTaskProcessor(t *testing.T) {
	size := 1_0000_1234
	vData := make([]any, 0, size)
	for i := 0; i < size; i++ {
		vData = append(vData, i)
	}

	processor := NewTaskProcessor(
		WithMaxWorkerCount(5),

		WithErrorHandler(func(err error) {
			log.Errorf("Error occurred: %v", err)
		}),

		WithResultHandler(func(data any) {
			log.Infof("%v", data)
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	err := processor.ProcessInChunks(ctx, vData,
		func(data []any) (any, error) {
			result := make([]any, 0, len(data))

			for i := range data {
				if val, ok := data[i].(int); ok {
					if val%2 == 0 {
						result = append(result, val)
					}
				}
			}

			return len(result), nil
		})
	if err != nil {
		log.Fatalf("Process failed: %v", err)
	}

	stats := processor.Metrics()
	log.Infof("Porcessing completed. Completed tasks: %d, Errors: %d", stats.completedTasks, stats.errorCount)
}

func TestTaskProcess_DataCut(t *testing.T) {
	size := 1234
	vData := make([]any, 0, size)
	for i := 0; i < size; i++ {
		vData = append(vData, i)
	}

	processor := NewTaskProcessor(
		WithMaxWorkerCount(runtime.GOMAXPROCS(0)),

		WithErrorHandler(func(err error) {
			log.Errorf("Error occurred: %v", err)
		}),

		WithResultHandler(func(data any) {
			log.Infof("%v", data)
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	err := processor.ProcessInChunks(ctx, vData,
		func(data []any) (any, error) {
			return len(data), nil
		})
	if err != nil {
		log.Fatalf("Process failed: %v", err)
	}

	stats := processor.Metrics()
	log.Infof(
		"Porcessing completed. Completed tasks: %d, Errors: %d",
		stats.completedTasks,
		stats.errorCount,
	)
}
