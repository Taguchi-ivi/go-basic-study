package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int) {
	defer wg.Done()
loop:
	// forの中でch1,ch2がnilになるまでloopする
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break loop
			// loopをコメントアウトして残っているバッファの値を取り出すこともできる。ただしバッファを指定していない場合はdeadlockになるので注意
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("ch1: %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("ch2: %v\n", v)
		}
	}
}

func main() {
	// select: receive continuous data
	// select文を使って複数のchannelの連続したデータを受信する
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	wg.Add(3)

	// goroutineでch1, ch2に値を入れる
	go countProducer(&wg, ch1, 5, 50)
	go countProducer(&wg, ch2, 5, 500)
	// goroutineでch1, ch2から値を受け取る
	go countConsumer(ctx, &wg, ch1, ch2)
	wg.Wait()
	fmt.Println("finish")

}
