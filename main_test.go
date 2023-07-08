package main

import (
	"testing"
)

func TestAddTest(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1+2=3",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "2+3=5",
			args: args{x: 2, y: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTest(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("AddTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivied(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "1/2=0.5",
			args: args{x: 1, y: 2},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divied(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Divied() = %v, want %v", got, tt.want)
			}
		})
	}
}
