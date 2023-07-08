package calculator

import "fmt"

// 小文字はprivate =>ただし同じパッケージ内からは呼び出せる
var offset float64 = 1

// 大文字rはpublic
var Offset float64 = 1

// 大文字で始めた関数はpublic
// 説明になる
func Sum(a float64, b float64) float64 {
	fmt.Println("mulutiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))
	return a + b + offset
}
