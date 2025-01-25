package main

import "testing"

func Test_findMedian(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Ex01", args: args{nums: []int{1, 5, 9}}, want: float64(5)},
		{name: "Ex01", args: args{nums: []int{1, 5, 9, 12}}, want: (float64(5) + float64(9)) / float64(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedian(tt.args.nums); got != tt.want {
				t.Errorf("findMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
