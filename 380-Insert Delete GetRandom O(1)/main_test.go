package main

import (
	"testing"
)

func TestRandomizedSet_Insert(t *testing.T) {
	type fields struct {
		m map[int]int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				m: map[int]int{},
			},
			args: args{
				val: 1,
			},
			want: true,
		}, {
			name: "",
			fields: fields{
				m: map[int]int{1: 1},
			},
			args: args{
				val: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedSet{
				m: tt.fields.m,
			}
			if got := this.Insert(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_Remove(t *testing.T) {
	type fields struct {
		m map[int]int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				m: map[int]int{1: 1},
			},
			args: args{
				val: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				m: map[int]int{1: 1},
			},
			args: args{
				val: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedSet{
				m: tt.fields.m,
			}
			if got := this.Remove(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_GetRandom(t *testing.T) {
	type fields struct {
		m map[int]int
		a []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				m: map[int]int{1: 1},
				a: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedSet{
				m: tt.fields.m,
				a: tt.fields.a,
			}
			if got := this.GetRandom(); got != tt.want {
				t.Errorf("RandomizedSet.GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
