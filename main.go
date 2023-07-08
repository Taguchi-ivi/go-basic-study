package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// import "fmt"
const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	i int
	s string
	b bool
)

// 構造体
type Task struct {
	Title    string
	Estimate int
}

func main() {

	// ### setup ###
	// fmt.Println("Hello World")
	// sl := []int{1, 2, 3}
	// if len(sl) > 2 {
	// 	fmt.Println("unreachable code")
	// }

	// ### module package ###
	// godotenv.Load()
	// fmt.Println(os.Getenv("GO_ENV"))
	// fmt.Println(calculator.Offset)
	// fmt.Println(calculator.Sum(1, 2))
	// fmt.Println(calculator.Multiply(1, 2))

	// ### variables ###
	// varで宣言すると関数外でも使える
	// :=で宣言すると関数内でしか使えない
	// i := 1
	// j := uint16(i)
	// fmt.Printf("i: %[1]v %[1]T j: %[2]v %[2]T\n", i, j)
	// pi, title := 3.14, "Go"
	// fmt.Printf("pi: %[1]v %[1]T title: %[2]v %[2]T\n", pi, title)
	// x := 1
	// y := 3.14
	// z := float64(x) + y
	// fmt.Println(z)
	// fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

	// ### point / shadowing ###
	// 2byteのうちの1byteを出力する &をつけることでメモリアドレス(先頭の1byte)を出力できる
	// 2byte分ずれているのは、1byte目にはデータが入っているが、2byte目には何も入っていないから
	// ポインタ変数はメモリ内のアドレスを格納する変数 ポインタ自体も8byteのメモリを使用する
	// shadowing :=で再定義しない(=のみ)だとshadowingにならず変数を更新できる
	// var ui1 uint16
	// fmt.Printf("memory address of ui1: %p\n", &ui1)
	// var ui2 uint16
	// fmt.Printf("memory address of ui2: %p\n", &ui2)
	// var p1 *uint16
	// fmt.Printf("value of p1: %v\n", p1)
	// p1 = &ui1
	// fmt.Printf("value of p1: %v\n", p1)
	// fmt.Printf("size of p1: %d[byte]\n", unsafe.Sizeof(p1))
	// fmt.Printf("memory address of p1: %p\n", &p1)
	// fmt.Printf("value of ui1(dereference): %v\n", *p1)
	// *p1 = 1
	// fmt.Printf("value of ui1: %v\n", ui1)
	// var pp1 **uint16 = &p1
	// fmt.Printf("value of pp1: %v\n", pp1)
	// fmt.Printf("memory address of pp1 %p\n", &pp1)
	// ok, result := true, "A"
	// if ok {
	// 	// result := "B"
	// 	result = "B"
	// 	fmt.Println(result)
	// } else {
	// 	// result := "C"
	// 	result = "C"
	// 	fmt.Println(result)
	// }
	// fmt.Println(result)

	// ### slice / map ###
	// var a1 [3]int
	// var a2 = [3]int{1, 2, 3}
	// 自動的に配列の要素数を指定してくれる ...
	// a3 := [...]int{10, 20}
	// fmt.Println(a1, a2, a3)
	// fmt.Printf("%v %v\n", len(a3), cap(a3))
	// sliceの定義 nilで判定できるとできないの違い
	// var s1 []int
	// s2 := []int{}
	// fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	// fmt.Println(s1 == nil, s2 == nil)
	// s1 = append(s1, 1, 2, 3)
	// fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// s3 := []int{4, 5, 6}
	// s1 = append(s1, s3...)
	// fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// // make 型と要素数とメモリを指定してスライスを作成する (要素数とcapを指定できる)
	// s4 := make([]int, 0, 2)
	// fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// s4 = append(s4, 1, 2, 3)
	// fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// s5 := make([]int, 4, 6)
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// s6 := s5[1:3]
	// s6[1] = 10
	// // sliceを切り取って使用するとメモリが共有される
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	// s6 = append(s6, 2)
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("s6 appended: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	// // メモリを共有しないようにするにはcopyを使う
	// sc6 := make([]int, len(s5[1:3]))
	// fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("sc6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// copy(sc6, s5[1:3])
	// fmt.Printf("sc6 dst copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// sc6[1] = 12
	// fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("sc6: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// // メモリを部分的に共有する場合は、ポインタを使う メモリを共有する最大数2(3-1)までメモリを共有する
	// s5 = make([]int, 4, 6)
	// fs6 := s5[1:3:3]
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	// fs6[0] = 6
	// fs6[1] = 7
	// fs6 = append(fs6, 8)
	// fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	// fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	// map [key]value nilと判定できるかの違い
	// var m1 map[string]int
	// m2 := map[string]int{}
	// fmt.Printf("%v %v \n", m1, m1 == nil)
	// fmt.Printf("%v %v \n", m2, m2 == nil)
	// m2["A"] = 10
	// m2["B"] = 20
	// m2["C"] = 0
	// fmt.Printf("%v %v %v \n", m2, len(m2), m2["A"])
	// delete(m2, "A")
	// fmt.Printf("%v %v %v \n", m2, len(m2), m2["A"])
	// // 要素が存在しないと0が返ってくる その判定
	// v, ok := m2["A"]
	// fmt.Println(v, ok)
	// v, ok = m2["C"]
	// fmt.Println(v, ok)
	// for k, v := range m2 {
	// 	fmt.Println(k, v)
	// }

	// ### struct / receiver ###
	// 上部で定義したstructを使用する
	// 代入しても値渡し(別々のメモリ領域)になる
	// ポインタにすることで関数でも値を変更できる
	// task1 := Task{
	// 	Title:    "Go",
	// 	Estimate: 60,
	// }
	// task1.Title = "Golang"
	// fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)
	// var task2 Task = task1
	// task2.Title = "Python"
	// fmt.Println(task1, task2)
	// task1p := &Task{
	// 	Title:    "concurrency",
	// 	Estimate: 120,
	// }
	// fmt.Println(task1p, *task1p)
	// task1p.Title = "Changed"
	// fmt.Println(task1p, *task1p)
	// var task2p *Task = task1p
	// task2p.Title = "changed by task2p"
	// fmt.Println("検証", task1p, *task1p, task2p, *task2p)
	// task1.extendEstimate()
	// fmt.Println("task1 value receiver", task1)
	// task1.extendEstimatePointer()
	// fmt.Println("task1 value Pointer", task1)

	// ### function / closure ###
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
	// 無名関数
	i := 1
	// 即座に実行するには最後に()をつける
	func(i int) {
		fmt.Println(i)
	}(i)
	// 即座に実行させない場合は、変数に代入してから実行する
	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))
	// 無名関数を関数の引数として渡す
	f2 := func(file string) string {
		return file + ".csv"
	}
	addExt(f2, "file1")
	// 無名関数をreturnで返すこともできる 使い道なかなかないから実装はしない
	// ここからclosure
	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Printf("%v\n", v)
	}
}

// receiver 型にメソッドを定義する(taskにextendEstimateを定義)
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// Pointer 型にメソッドを定義する(taskにextendEstimateを定義)
func (task *Task) extendEstimatePointer() {
	task.Estimate += 10
}

func funcDefer() {
	// deferは関数の最後に実行される
	// ファイルをcloseさせるときに使う。ミス防止
	// 複数のdefer文がある場合は、最後に定義したもの(下から順に)から実行される
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("hello world")
}

// 可変な引数に対応する
func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

// 複数の戻り値を返す場合の指定
func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
}

// countUp内で宣言したcountの変数はcountUp内でglobalな値となり、countUpを呼び出すたびに値が保持される(閉じ込められている)
func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}
