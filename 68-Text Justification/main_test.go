package main

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Ex01", args: args{words: []string{"This", "is", "an", "example", "of", "text", "justification."}, maxWidth: 16}, want: []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{name: "Ex02", args: args{words: []string{"What", "must", "be", "acknowledgment", "shall", "be"}, maxWidth: 16}, want: []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{name: "Ex03", args: args{words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, maxWidth: 20}, want: []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
