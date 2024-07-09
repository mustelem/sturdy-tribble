package main

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	r := RandomizedSet{}
	r.m = make(map[int]int)
	r.a = make([]int, 0, 20000)

	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.m[val]

	if !exists {
		this.a = append(this.a, val)
		this.m[val] = val
	}

	return !exists
}

func (this *RandomizedSet) Remove(val int) bool {
	_, exists := this.m[val]

	if exists {
		delete(this.m, val)
	}

	return exists
}

func (this *RandomizedSet) GetRandom() int {
	length := len(this.a)

	for {
		idx := rand.Intn(length)

		_, exists := this.m[this.a[idx]]
		if exists {
			return this.a[idx]
		}
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
