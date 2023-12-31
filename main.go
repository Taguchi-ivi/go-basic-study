package main

import (
	"context"
	"fmt"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
)

// 並列処理
// 各スレッドでの処理を同時に行うこと

// 並行処理
// 1つのスレッド内で複数の処理を切り替えながら行うこと
// Go言語は並行処理を行うことができる

// 作成されたtrace.outを開く
// go tool trace trace.out

// データが競合(データレース)しているか確認する
// go run -race main.go

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
	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	log.Fatalln("Error:", err)
	// }
	// defer func() {
	// 	// if文の中でerrを宣言して条件分岐として使うこともできる
	// 	if err := f.Close(); err != nil {
	// 		log.Fatalln("Error:", err)
	// 	}
	// }()
	// // トレースの開始
	// if err := trace.Start(f); err != nil {
	// 	log.Fatalln("Error:", err)
	// }
	// defer trace.Stop()
	// // トレースの実際の処理 mainという名前でトレースを開始する
	// ctx, t := trace.NewTask(context.Background(), "main")
	// defer t.End()
	// fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())
	// // 逐次処理の場合
	// // task(ctx, "Task1")
	// // task(ctx, "Task2")
	// // task(ctx, "Task3")
	// // 並列処理の場合
	// var wg sync.WaitGroup
	// wg.Add(3)
	// go cTask(ctx, &wg, "Task1")
	// go cTask(ctx, &wg, "Task2")
	// go cTask(ctx, &wg, "Task3")
	// s := []int{1, 2, 3}
	// // goroutineの中だと起動に少し時間がかかりiの値が更新されて出力されてしまう。
	// // そのため、goroutineの中でiを使う場合は引数として渡す必要がある
	// // goroutineでは完全に別のスレッドで実行されるため、順番が前後することがある
	// for _, i := range s {
	// 	wg.Add(1)
	// 	// go func() {
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}(i)
	// 	// }()
	// }
	// wg.Wait()
	// fmt.Println("main func finish")

	// ### goroutine:channel ###
	// チャネルを使うことによって、goroutine間でデータのやり取りを行うことができる
	// チャネルへの書き込み操作は、チャネルへの読み込み操作が行われるまでブロックされる(バッファがない場合)
	// やりとりするデータの型を指定する必要がある 下記ではint型
	// <- はチャネルの読み込み操作 / / ch <- はチャネルへの書き込み操作
	// 別のgoroutineの値をチャネル経由で受け取る
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
	// fmt.Println("main func finish")
	// goroutineリーク 受信側が送信されるまで待ち続けるため、メモリが解放されないことを言う
	// 以下のようにgoroutineを作成すると、受信側が送信されるまで待ち続けるため、goroutineが終了しない
	// ch1 := make(chan int)
	// go func() {
	// 	fmt.Println(<-ch1)
	// }()
	// // 下記の書き込みの処理がないとgoroutineリークが起きる
	// ch1 <- 10
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// // バッファを指定することで、goroutineリークを防ぐことができる
	// // バッファを指定すると、バッファの分だけ書き込みができる(バッファの分だけ読み込みを待たずに書き込みができる)
	// // バッファを超える書き込み or 読み込みから書き込みの順番になると deadlock が起きる
	// ch2 := make(chan int, 1)
	// ch2 <- 2
	// // ch2 <- 3
	// fmt.Println(<-ch2)

	// ### channel close, capsel, notification ###
	// チャネルを閉じることで、チャネルに対する書き込みを禁止することができる
	// ただしバッファがある場合は、バッファにある値を読み込み終わるまで読み込みを許可する その後は読み込みを禁止する
	// ch1 := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println(<-ch1)
	// }()
	// ch1 <- 10
	// close(ch1)
	// v, ok := <-ch1
	// fmt.Println(v, ok)
	// wg.Wait()
	// ch2 := make(chan int, 2)
	// ch2 <- 1
	// ch2 <- 2
	// close(ch2)
	// v, ok = <-ch2
	// fmt.Println(v, ok)
	// v, ok = <-ch2
	// fmt.Println(v, ok)
	// v, ok = <-ch2
	// fmt.Println(v, ok)

	// ch3 := generateCountStream()
	// for v := range ch3 {
	// 	fmt.Println(v)
	// }
	// // データの値を持たない通知専用のチャネル
	// // closeすると deadlockが解除されて print処理がはじまる
	// nCh := make(chan struct{})
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Printf("goroutine %v started\n", i)
	// 		<-nCh
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// time.Sleep(2 * time.Second)
	// close(nCh)
	// fmt.Println("unblocked by manual close")
	// wg.Wait()
	// fmt.Println("finish")

	// ### select with timeout context default ###
	// selectを使うことで、複数のチャネルの受信ができるようになる
	// 	const byfSize = 3
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
	// 		ch2 <- "B"
	// 	}()
	// 	// どちらかの値がnilになった時ループを抜ける
	// loop:
	// 	for ch1 != nil || ch2 != nil {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("timeout")
	// 			break loop
	// 		case v := <-ch1:
	// 			fmt.Println(v)
	// 			ch1 = nil
	// 		case v := <-ch2:
	// 			fmt.Println(v)
	// 			ch2 = nil
	// 		default:
	// 			fmt.Println("no msg arrived")
	// 		}
	// 	}
	// 	wg.Wait()
	// 	fmt.Println("finish")

	// ### select receive continuous data ###
	// セレクト文を用いて、複数のチャネルから連続的にデータを受信する
	// const bufSize = 5
	// ch1 := make(chan int, bufSize)
	// ch2 := make(chan int, bufSize)
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	// defer cancel()
	// wg.Add(3)
	// go countProducer(&wg, ch1, bufSize, 50)
	// go countProducer(&wg, ch2, bufSize, 500)
	// go countConsumer(ctx, &wg, ch1, ch2)
	// wg.Wait()
	// fmt.Println("finish")

	// ### mutex + atomic ###
	// データの競合(データレース)を防ぐために、排他制御を行う
	// var wg sync.WaitGroup
	// var i int
	// var mu sync.Mutex
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	i++
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	i++
	// }()
	// wg.Wait()
	// fmt.Println(i)
	// rwMutex
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
	// atomic
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

// カプセル化
// 読み込み専用のチャネルを生産するための関数
// 読み取り専用のチャネルだけにアクセスすることで、チャネルの生成、書き込み、クローズをカプセル化することができる
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}

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
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break loop
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("ch1", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Println("ch2", v)
		}
	}
}

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
