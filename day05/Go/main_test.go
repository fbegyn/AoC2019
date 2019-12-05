package main

import (
	"reflect"
	"testing"
)

func Test_runInstruction(t *testing.T) {

	example1 := []int{1, 0, 0, 0, 99}
	example2 := []int{1, 0, 0, 0, 93}
	example3 := []int{1002, 4, 3, 4, 33}
	example4 := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	type args struct {
		program *[]int
		index   int
	}
	tests := []struct {
		name string
		args args
		run  bool
		jump int
	}{
		{"jump sum", args{&example1, 0}, true, 4},
		{"jump halt", args{&example1, 4}, false, 1},
		{"jump non opcode", args{&example2, 4}, false, 0},
		{"jump multiplication", args{&example3, 0}, true, 4},
		{"jump if false", args{&example4, 2}, true, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run, jump := runInstruction(tt.args.program, tt.args.index)
			if run != tt.run {
				t.Errorf("runInstruction() = %v, want %v", run, tt.run)
			}
			if jump != tt.jump {
				t.Errorf("runInstruction() = %v, want %v", jump, tt.jump)
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
		{"test 5", args{[]int{1002, 4, 3, 4, 33}}, []int{1002, 4, 3, 4, 99}},
		{"test 6", args{[]int{1101, 100, -1, 4, 0}}, []int{1101, 100, -1, 4, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runProgram(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
