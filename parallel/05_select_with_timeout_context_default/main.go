package main

import (
	"sync"
	"time"
)

func main() {
	// select with timeout and context
	// 複数のチャンネルからの受信を待つ場合、select文を使う

	// not use context
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(500 * time.Millisecond)
	// 	ch1 <- "A"
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(800 * time.Millisecond)
	// 	ch1 <- "B"
	// }()
	// for ch1 != nil || ch2 != nil {
	// 	select {
	// 	case v := <-ch1:
	// 		fmt.Println(v)
	// 		ch1 = nil
	// 	case v := <-ch2:
	// 		fmt.Println(v)
	// 		ch2 = nil
	// 	}
	// }
	// wg.Wait()
	// fmt.Println("finish")

	// use context
	// bufferをつけないとdeadlockする timeoutになった際にch1, ch2に値が入らないため
	// 	ch1 := make(chan string, 1)
	// 	ch2 := make(chan string, 1)
	// 	var wg sync.WaitGroup
	// 	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	// 	defer cancel()
	// 	wg.Add(2)
	// 	go func() {
	// 		defer wg.Done()
	// 		time.Sleep(500 * time.Millisecond)
	// 		ch1 <- "A"
	// 	}()
	// 	go func() {
	// 		defer wg.Done()
	// 		time.Sleep(800 * time.Millisecond)
	// 		ch1 <- "B"
	// 	}()

	// loop:
	// 	for ch1 != nil || ch2 != nil {
	// 		select {
	// 		case <-ctx.Done():
	// 			// contextがキャンセルされた場合
	// 			fmt.Println("timeout")
	// 			break loop
	// 		case v := <-ch1:
	// 			fmt.Println(v)
	// 			ch1 = nil
	// 		case v := <-ch2:
	// 			fmt.Println(v)
	// 			ch2 = nil
	// 		}
	// 	}
	// 	wg.Wait()
	// 	fmt.Println("finish")

	// select default
	var wg sync.WaitGroup
	ch := make(chan string, 3)
	wg.Add(1)
	// 書き込み側のgoroutine
	go func() {
		defer wg.Done()
		ch <- "A"
		for i := 0; i < 3; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- "hello"
		}
	}()
	// 読み込み側のgoroutine
	for i := 0; i < 3; i++ {
		select {
		case v := <-ch:
			println(v)
		default:
			println("no msg arrived")
		}
		time.Sleep(1500 * time.Millisecond)
	}
}
