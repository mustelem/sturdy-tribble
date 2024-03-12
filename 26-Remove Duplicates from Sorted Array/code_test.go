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
		{
			name: "two same vals",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "two vals",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "happy",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "empty array",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{name: "some",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
