package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Ex 1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "Ex 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
		{
			name: "Edge 1",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
