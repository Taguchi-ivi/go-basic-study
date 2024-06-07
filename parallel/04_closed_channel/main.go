package main

import (
	"fmt"
	"sync"
)

// カプセル化 <-読み取り専用
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	// channel close, capsel, notification

	// channel close
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	close(ch1)
	// v: 10, ok: true / v: 0, ok: false 読み込む値がない場合は0とfalseが返る
	v, ok := <-ch1
	fmt.Printf("v: %d, ok: %v\n", v, ok)
	wg.Wait()

	// buffer channel
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	v, ok = <-ch2
	fmt.Printf("v: %d, ok: %v\n", v, ok) // v: 1, ok: true
	v, ok = <-ch2
	fmt.Printf("v: %d, ok: %v\n", v, ok) // v: 2, ok: true
	v, ok = <-ch2
	fmt.Printf("v: %d, ok: %v\n", v, ok) // v: 0, ok: false bufferの場合はcloseしてもbufferの値が全て読み込まれるまでokはtrue

	// capsel, 読み取り専用の関数を作成
	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
	}

	// notification
	// struct{}型は0バイトしか消費されない、通知のchannelには適している
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %d\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")

}
