package main

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
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
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "Ex 2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
		{
			name: "failed test 1",
			args: args{
				ratings: []int{1, 2, 87, 87, 87, 2, 1},
			},
			want: 13,
		},
		{
			name: "failed test 2",
			args: args{
				ratings: []int{29, 51, 87, 87, 72, 12},
			},
			want: 12,
		},
		{
			name: "failed test 3",
			args: args{
				ratings: []int{1, 3, 4, 5, 2},
			},
			want: 11,
		},
		{
			name: "my test 1",
			args: args{
				ratings: []int{1, 5, 4, 3, 2},
			},
			want: 11,
		},
		{
			name: "my test 2",
			args: args{
				ratings: []int{2, 3, 4, 5, 2, 3},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
