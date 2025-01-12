package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Ex01", args: args{num: 3749}, want: "MMMDCCXLIX"},
		{name: "Ex02", args: args{num: 58}, want: "LVIII"},
		{name: "Ex03", args: args{num: 3}, want: "III"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
