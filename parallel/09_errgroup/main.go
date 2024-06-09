package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

//	func doTask(task string) error {
//		if task == "fake1" || task == "fake2" {
//			return fmt.Errorf("%v failed", task)
//		}
//		fmt.Printf("task %v completed\n", task)
//		return nil
//	}
func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "fake1":
		t = time.NewTicker(500 * time.Millisecond)
	case "fake2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}
	select {
	case <-ctx.Done():
		fmt.Printf("%v canceled: %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		if task == "fake1" || task == "fake2" {
			return fmt.Errorf("%v failed", task)
		}
		fmt.Printf("task %v completed\n", task)
	}
	return nil
}

func main() {
	// errGroup
	// 複数のgoroutineのエラーを集約して扱う
	// eg.Go(func() error のようにgoroutineを起動することでerrorを受け取ることができる
	// contextと組み合わせることで、一個でもエラーがあったら全てのgoroutineをキャンセルすることもできる
	// 外部パッケージ
	// eg := new(errgroup.Group)
	eg, ctx := errgroup.WithContext(context.Background())
	s := []string{"task1", "fake1", "task2", "fake2", "task3"}
	for _, v := range s {
		task := v
		eg.Go(func() error {
			// return doTask(task)
			return doTask(ctx, task)
		})
	}
	// Wait()で全てのgoroutineが終了するまで待つ
	if err := eg.Wait(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("finish")
}
