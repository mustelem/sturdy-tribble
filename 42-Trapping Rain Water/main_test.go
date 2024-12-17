package main

import (
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 01",
			args: args{
				height: []int{1, 0, 1},
			},
			want: 1,
		}, {
			name: "Test 02",
			args: args{
				height: []int{1, 0, 2},
			},
			want: 1,
		}, {
			name: "Test 03",
			args: args{
				height: []int{2, 0, 1},
			},
			want: 1,
		}, {
			name: "Test 04",
			args: args{
				height: []int{2, 0, 2},
			},
			want: 2,
		}, {
			name: "Test 05",
			args: args{
				height: []int{2, 0, 1, 2},
			},
			want: 3,
		}, {
			name: "Test 06",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		}, {
			name: "Test 07",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		}, {
			name: "Test 08",
			args: args{
				height: []int{1, 3, 2, 1, 2, 1, 2, 1, 3},
			},
			want: 9,
		}, {
			name: "Test 09",
			args: args{
				height: []int{5, 4, 1, 2},
			},
			want: 1,
		}, {
			name: "Test 10",
			args: args{
				height: []int{5, 4, 1, 1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
