package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
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
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "Ex 2",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		},
		{
			name: "Ex 3",
			args: args{
				s: "luffy is still joyboy",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
