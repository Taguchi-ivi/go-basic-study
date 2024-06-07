package main

import (
	"fmt"
	"runtime"
)

func main() {
	// channel unbuffer / buffered / goroutine leak
	// <-read(読み込み) / write(書き込み)<-
	// <- channel <-
	// fmt.Println(<-ch) // read
	// ch <- 1 // write

	// unbuffered channel
	// 別のgoroutineで書き込んだ値をmain goroutineで読み込む
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// goroutine leak
	// goroutineリークが起こっていないか確認するライブラリ go.uber.org/goleak
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	// 下記一行がないとgoroutineリークが発生する
	ch1 <- 10 // goroutine leak
	fmt.Printf("num of waiting goroutines: %d\n", runtime.NumGoroutine())

	// buffered channel
	// bufferがいっぱいになるまで書き込みが可能(本来は読み込みがないと書き込みはできない)
	ch2 := make(chan int, 1)
	ch2 <- 2
	fmt.Println(<-ch2)
}
