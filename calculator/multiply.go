package calculator

func Multiply(a float64, b float64) float64 {
	return (a * b) + offset
}

// 小文字で始めた関数はprivate => ただし同じパッケージ内からは呼び出せる
func multiply(a float64, b float64) float64 {
	return (a * b) + offset
}
