package main

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test one",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		}, {
			name: "test two",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		}, {
			name: "needle too big",
			args: args{
				haystack: "recep",
				needle:   "mehmet",
			},
			want: -1,
		},
		{
			name: "bu ne",
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
