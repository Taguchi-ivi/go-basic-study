package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func subTaskForWithTimeout(ctx context.Context, wg *sync.WaitGroup, id string) {
	defer wg.Done()
	t := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-t.C:
			t.Stop()
			fmt.Println(id)
			return
		}
	}
}

func criticalTask(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()
	t := time.NewTicker(1000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "A", nil
}

func normalTask(ctx context.Context) (string, error) {
	t := time.NewTicker(3000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "B", nil
}

func subTaskForDeadline(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		// deadlineの時刻を取得, 設定の有無でokがtrueかfalseで返ってくる
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "hello"
	}()
	return ch

}

func main() {
	// context: cancel, timeout, deadline

	// timeoutを使うことによって、指定時間後に全てのgoroutineをキャンセルする
	// var wg sync.WaitGroup
	// // 600ms後にキャンセル
	// ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	// defer cancel()
	// go subTaskForWithTimeout(ctx, &wg, "a")
	// go subTaskForWithTimeout(ctx, &wg, "b")
	// go subTaskForWithTimeout(ctx, &wg, "c")
	// wg.Wait()

	// cancelを使うことによって、任意のタイミングで全てのgoroutineをキャンセルする
	// critical taskはtimeoutによってキャンセルされる, それによってcancel()が呼ばれるので、normal taskもキャンセルされる
	// err内容: critical task cancelled due to context deadline exceeded
	// err内容: normal task cancelled due to context canceled
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := criticalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("critical task cancelled due to %v\n", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("success:", v)
	// }()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := normalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("normal task cancelled due to %v\n", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("success:", v)
	// }()
	// wg.Wait()

	// deadlineを使うことによって、指定時間後に全てのgoroutineをキャンセルする
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(40*time.Millisecond))
	defer cancel()
	ch := subTaskForDeadline(ctx)
	// 閉じているか判定
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")
}
