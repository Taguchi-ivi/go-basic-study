package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")
}
func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("write lock")
	*c++
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")
}

func main() {
	// mutex, atomic
	// 排他制御をしないとデータ競合が発生する
	// データ競合の確認: go run -race main.go
	// var wg sync.WaitGroup
	// var i int
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	i++
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	i++
	// }()
	// wg.Wait()
	// fmt.Println(i) // 1 or 2

	// mutexを使って排他制御を行う
	// 再度 go run -race main.go をしてもデータ競合が発生しない
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// var i int
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	i++
	// 	mu.Unlock()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	i++
	// 	mu.Unlock()
	// }()
	// wg.Wait()
	// fmt.Println(i) // 2

	// rwmutexを使って読み込みと書き込みを分ける
	// read時はlockしていても資源を共有するため、複数のgoroutineが同時に読み込みを行うことができる
	// var wg sync.WaitGroup
	// var rwMu sync.RWMutex
	// var c int

	// wg.Add(4)
	// go write(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// wg.Wait()
	// fmt.Println("finish")

	// atomicを使って排他制御を行う
	// atomicを使うとmutexを使うよりも高速(軽量)に処理ができる
	// ただし、atomicはint型のみ対応している
	// また、atomicはadd, compareAndSwap, load, storeなどの関数がある
	// atomicは1つの操作についてのみ排他制御ができる
	var wg sync.WaitGroup
	var c int64

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&c, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(c)
	fmt.Println("finish")
}
