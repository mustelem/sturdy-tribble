package main

import "testing"

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
			name: "Should hold 1 unit of water",
			args: args{
				height: []int{1, 0, 1},
			},
			want: 1,
		}, {
			name: "Should hold 1 unit of water",
			args: args{
				height: []int{1, 0, 2},
			},
			want: 1,
		}, {
			name: "Should hold 1 unit of water",
			args: args{
				height: []int{2, 0, 1},
			},
			want: 1,
		}, {
			name: "Should hold 2 unit of water",
			args: args{
				height: []int{2, 0, 2},
			},
			want: 2,
		}, {
			name: "Should hold 3 unit of water",
			args: args{
				height: []int{2, 0, 1, 2},
			},
			want: 3,
		}, {
			name: "Should hold 6 unit of water",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		}, {
			name: "Should hold 9 unit of water",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
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
