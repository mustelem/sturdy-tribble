package main

func trap(height []int) int {
	var tempAcc, permAcc int
	var maxHeight int

	for _, v := range height {
		if v < maxHeight {
			tempAcc += maxHeight - v
		} else if v > 0 {

		} else {
			permAcc += tempAcc
			tempAcc = 0
		}
		maxHeight = max(v, maxHeight)
	}

	return 0
}
