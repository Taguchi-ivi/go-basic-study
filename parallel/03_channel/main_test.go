package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	// "go.uber.org/goleak" はgoroutineリークが起こっていないか確認するライブラリ
	defer goleak.VerifyNone(t)
	main()
}
