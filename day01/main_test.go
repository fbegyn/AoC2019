package main

import (
	"testing"
)

func Test_calcFuel(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name  string
		args  args
		wantF float64
	}{
		{"simple", args{12}, 2},
		{"average", args{1969}, 966},
		{"complex", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotF := calcFuel(tt.args.mass); gotF != tt.wantF {
				t.Errorf("calcFuel() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name      string
		args      args
		wantLevel int
	}{
		{"test.txt input", args{"test.txt"}, 34241},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLevel := part1(tt.args.file); gotLevel != tt.wantLevel {
				t.Errorf("part1() = %v, want %v", gotLevel, tt.wantLevel)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name      string
		args      args
		wantLevel int
	}{
		{"test.txt input", args{"test.txt"}, 51316},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLevel := part2(tt.args.file); gotLevel != tt.wantLevel {
				t.Errorf("part2() = %v, want %v", gotLevel, tt.wantLevel)
			}
		})
	}
}
