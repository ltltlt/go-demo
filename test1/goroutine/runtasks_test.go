package runtasks_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func RunTasks(ctx context.Context, concurrency int, tasks ...func() error) error {
	ch := make(chan func() error)
	errChan := make(chan error, 1)

	// producer
	var err error
	go func() {
		for _, task := range tasks {
			select {
			case ch <- task:
			case err = <-errChan:
				break
			case <-ctx.Done():
				err = ctx.Err()
				break
			}
		}
		close(ch)
	}()

	// consumer
	count := concurrency
	if count > len(tasks) {
		count = len(tasks)
	}
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					select {
					case errChan <- err:
					default:
					}
				}
			}
		}()
	}

	wg.Wait()

	if err != nil {
		return err
	}

	select {
	case err = <-errChan:
		return err
	default:
	}

	return nil
}

func TestRunTasks(t *testing.T) {
	var err = errors.New("abc")
	var err1 = errors.New("e")
	tasks := []func() error{
		func() error {
			return err1
		},
		func() error {
			time.Sleep(time.Second * 2)
			return err
		}, func() error {
			time.Sleep(time.Second)
			return nil
		},
	}

	err2 := RunTasks(context.Background(), 100, tasks...)
	assert.Equal(t, err1, err2)
}
