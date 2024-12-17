package main

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Ex 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		}, {
			name: "Ex 2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		}, {
			name: "Ex 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		}, {
			name: "Edge 1",
			args: args{
				ransomNote: "",
				magazine:   "aab",
			},
			want: true,
		}, {
			name: "Edge 2",
			args: args{
				ransomNote: "aa",
				magazine:   "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
