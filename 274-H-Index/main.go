package main

func hIndex(citations []int) int {
	counts := make([]int, 5000)

	for _, v := range citations {
		counts[v]++
	}

	var papers int
	for i := len(counts) - 1; i >= 0; i-- {
		papers += counts[i]
		if papers >= i {
			return i
		}
	}

	return 0
}
