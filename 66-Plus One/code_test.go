package main

import (
	"reflect"
	"testing"
)

func Test_toNum(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "happy",
			args: args{
				digits: []int{9, 3, 4},
			},
			want: 934,
		}, {
			name: "single digit",
			args: args{
				digits: []int{5},
			},
			want: 5,
		}, {
			name: "large number",
			args: args{
				digits: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			},
			want: 728509129536673284379577474947011174006,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toNum(tt.args.digits); got != tt.want {
				t.Errorf("toNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toArray(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "475",
			args: args{
				number: 475,
			},
			want: []int{4, 7, 5},
		}, {
			name: "5",
			args: args{
				number: 5,
			},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toArray(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "two digits affected",
			args: args{
				digits: []int{5, 8, 9},
			},
			want: []int{5, 9, 0},
		},
		{
			name: "all digits affected",
			args: args{
				digits: []int{9, 9, 9, 9, 9, 9},
			},
			want: []int{1, 0, 0, 0, 0, 0, 0},
		}, {
			name: "failed test",
			args: args{
				digits: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			},
			want: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
