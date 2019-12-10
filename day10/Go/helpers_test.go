package main

import (
	"testing"
)

func TestSumOfFloat64Array(t *testing.T) {
	type args struct {
		test []float64
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{"single", args{[]float64{1.0}}, 1.0},
		{"2 values", args{[]float64{1.0, 10.0}}, 11.0},
		{"multiple values", args{[]float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}}, 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SumOfFloat64Array(tt.args.test); gotResult != tt.wantResult {
				t.Errorf("SumOfFloat64Array() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSumOfIntArray(t *testing.T) {
	type args struct {
		test []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"single", args{[]int{1}}, 1},
		{"2 values", args{[]int{1, 10}}, 11},
		{"multiple values", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SumOfIntArray(tt.args.test); gotResult != tt.wantResult {
				t.Errorf("SumOfIntArray() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"little distance", args{[]int{0, 0}, []int{1, 1}}, 2},
		{"slightly larger distance distance", args{[]int{-5, -5}, []int{5, 5}}, 20},
		{"large boxy distance", args{[]int{-2500, -500}, []int{5000, 750}}, 8750},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManhattanDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"positive", args{3000}, 3000},
		{"negative", args{-9000}, 9000},
		{"zero", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
