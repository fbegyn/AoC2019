package main

import (
	"reflect"
	"testing"
)

func Test_checkPassword(t *testing.T) {
	type args struct {
		password []int
	}
	tests := []struct {
		name      string
		args      args
		wantPart1 bool
		wantPart2 bool
	}{
		{"Password 1", args{[]int{1, 1, 1, 1, 1, 1}}, false, true},
		{"Password 2", args{[]int{2, 2, 3, 4, 5, 0}}, true, true},
		{"Password 3", args{[]int{1, 2, 3, 7, 8, 9}}, false, false},
		{"Password 4", args{[]int{1, 1, 2, 2, 3, 3}}, true, true},
		{"Password 5", args{[]int{1, 2, 3, 4, 4, 4}}, false, true},
		{"Password 6", args{[]int{1, 1, 1, 1, 2, 2}}, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPart1, gotPart2 := checkPassword(tt.args.password)
			if gotPart1 != tt.wantPart1 {
				t.Errorf("checkPassword() gotPart1 = %v, want %v", gotPart1, tt.wantPart1)
			}
			if gotPart2 != tt.wantPart2 {
				t.Errorf("checkPassword() gotPart2 = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}

func Test_isIncreasing(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
		want2 int
	}{
		{"Low input", args{10}, false, []int{0, 0, 0, 0, 1, 0}, 11},
		{"Lower limit", args{193651}, false, []int{1, 9, 3, 6, 5, 1}, 199000},
		{"Lower limit next", args{199000}, false, []int{1, 9, 9, 0, 0, 0}, 199900},
		{"About the middle", args{367801}, false, []int{3, 6, 7, 8, 0, 1}, 367880},
		{"About the middle next", args{367880}, false, []int{3, 6, 7, 8, 8, 0}, 367888},
		{"Upper limit", args{649729}, false, []int{6, 4, 9, 7, 2, 9}, 660000},
		{"Upper limit next", args{660000}, false, []int{6, 6, 0, 0, 0, 0}, 666000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := isIncreasing(tt.args.n)
			if got != tt.want {
				t.Errorf("isIncreasing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("isIncreasing() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("isIncreasing() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
