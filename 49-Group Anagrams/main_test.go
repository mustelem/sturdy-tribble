package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "Ex01", args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{name: "Ex02", args: args{strs: []string{""}}, want: [][]string{{""}}},
		{name: "Ex03", args: args{strs: []string{"a"}}, want: [][]string{{"a"}}},
		{name: "Ex04", args: args{strs: []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}}, want: [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
