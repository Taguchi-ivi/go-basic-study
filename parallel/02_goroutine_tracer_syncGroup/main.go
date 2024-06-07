package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func main() {
	// goroutine: tracer+syncGroup
	// goroutineの中だと起動に少し時間がかかり、mainの処理が終わってしまう
	// fork join model
	// go func() {
	// 	fmt.Println("goroutine invoked")
	// }()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine()) // 2
	// fmt.Println("main func finish")

	// waitGroup
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("goroutine invoked")
	// }()
	// // waitGroupが終了するのを待つ
	// wg.Wait()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine()) // 1
	// fmt.Println("main func finish")

	// tracer
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// 作成されたtrace.outを開くコマンド
	// go tool trace trace.out

	// 逐次処理の場合、1秒ずつ処理に時間がかかる
	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")

	// 並列処理の場合、3つの処理が同時に実行される
	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "task1")
	go cTask(ctx, &wg, "task2")
	go cTask(ctx, &wg, "task3")
	wg.Wait()
	fmt.Println("main func finish")
}
