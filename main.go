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

// 並列処理
// 各スレッドでの処理を同時に行うこと

// 並行処理
// 1つのスレッド内で複数の処理を切り替えながら行うこと
// Go言語は並行処理を行うことができる

// 作成されたtrace.outを開く
// go tool trace trace.out

func main() {

	// ### goroutine:tracer + syncGroup ###
	// 先頭をgoにすることでgoroutineとして実行される
	// mainとは別のスレッドで実行される
	// 本来スレッドはmainで呼び出されても少し遅れて実行されるので、先にmainの処理が終了してしまう/=>そのため全てを一緒に実行するためにsyncGroupを使う
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("goroutine invoked")
	// }()
	// // ここでgoroutineが終了するのを待つ
	// wg.Wait()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Println("main func finish")
	// #### tracer ####
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		// if文の中でerrを宣言して条件分岐として使うこともできる
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
	// トレースの開始
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()
	// トレースの実際の処理 mainという名前でトレースを開始する
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())
	// 逐次処理の場合
	// task(ctx, "Task1")
	// task(ctx, "Task2")
	// task(ctx, "Task3")
	// 並列処理の場合
	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	s := []int{1, 2, 3}
	// goroutineの中だと起動に少し時間がかかりiの値が更新されて出力されてしまう。
	// そのため、goroutineの中でiを使う場合は引数として渡す必要がある
	// goroutineでは完全に別のスレッドで実行されるため、順番が前後することがある
	for _, i := range s {
		wg.Add(1)
		// go func() {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
		// }()
	}
	wg.Wait()
	fmt.Println("main func finish")
}

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
