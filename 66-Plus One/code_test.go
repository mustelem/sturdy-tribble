package main

import (
	"reflect"
	"testing"
)

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
			name: "",
			args: args{
				digits: []int{2, 8, 4},
			},
			want: []int{2, 8, 5},
		},
		{
			name: "",
			args: args{
				digits: []int{2, 8, 9},
			},
			want: []int{2, 9, 0},
		},
		{
			name: "",
			args: args{
				digits: []int{9, 9},
			},
			want: []int{1, 0, 0},
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
