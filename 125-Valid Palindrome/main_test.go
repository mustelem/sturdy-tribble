package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
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
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "Ex 2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "Ex 3",
			args: args{
				s: " ",
			},
			want: true,
		},
		{
			name: "Ex 4",
			args: args{
				s: "aa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlphaNumeric(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "non alpha numeric",
			args: args{
				r: 0,
			},
			want: false,
		},
		{
			name: "first letter",
			args: args{
				r: 97,
			},
			want: true,
		},
		{
			name: "last letter",
			args: args{
				r: 122,
			},
			want: true,
		},
		{
			name: "first number",
			args: args{
				r: 48,
			},
			want: true,
		},
		{
			name: "last number",
			args: args{
				r: 57,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlphaNumeric(tt.args.r); got != tt.want {
				t.Errorf("isAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
