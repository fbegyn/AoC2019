package main

import (
	"reflect"
	"testing"
)

func Test_runInstruction(t *testing.T) {

	example1 := []int{1, 0, 0, 0, 99}
	example2 := []int{1, 0, 0, 0, 93}

	type args struct {
		program *[]int
		index   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test false", args{&example1, 0}, true},
		{"test false", args{&example1, 4}, false},
		{"test true", args{&example2, 0}, true},
		{"test true", args{&example2, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runInstruction(tt.args.program, tt.args.index); got != tt.want {
				t.Errorf("runInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setParam(t *testing.T) {

	example := []int{1, 0, 0, 0, 99}

	type args struct {
		noun    int
		verb    int
		program []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 12 2", args{12, 2, example}, []int{1, 12, 2, 0, 99}},
		{"test 24 5", args{24, 5, example}, []int{1, 24, 5, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setParam(tt.args.noun, tt.args.verb, tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runProgram(t *testing.T) {
	type args struct {
		program []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test 1", args{[]int{1, 0, 0, 0, 99}}, []int{2, 0, 0, 0, 99}},
		{"test 2", args{[]int{2, 3, 0, 3, 99}}, []int{2, 3, 0, 6, 99}},
		{"test 3", args{[]int{2, 4, 4, 5, 99, 0}}, []int{2, 4, 4, 5, 99, 9801}},
		{"test 4", args{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runProgram(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
