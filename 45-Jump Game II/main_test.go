package main

import "testing"

func Test_jump(t *testing.T) {
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
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		}, {
			name: "Ex 2",
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
		{
			name: "failed test 1",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "my test 1",
			args: args{
				nums: []int{2, 4, 2, 8, 3, 4, 1, 3, 2, 1, 0, 4},
			},
			want: 3,
		},
		{
			name: "my test 2",
			args: args{
				nums: []int{2, 4, 2, 8, 3, 4, 1, 3, 2, 1, 1, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
