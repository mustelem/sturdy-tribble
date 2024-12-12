package main

func trap(height []int) int {
	lWall, rWall, trappedWater := false, false, 0

	var l, r int

	var temp []int = make([]int, len(height))

	var left int = 0
	for i := 0; i < len(height); i++ {
		h := height[i]

		if i > 0 {
			left = i - 1
		}

		if lWall && h > height[left] {
			r = i
			rWall = true
		}

		if lWall {
			temp[i] = height[l] - h
		}

		if lWall && rWall && -l > 1 {
			minH := min(height[l], height[r])

			for ptr := l + 1; ptr < r; ptr++ {
				canHoldAdditionally := min(minH-height[ptr], temp[ptr])
				trappedWater += canHoldAdditionally
				temp[ptr] -= canHoldAdditionally
			}
		}

		// and no temp water
		if lWall {
			if h > 0 {
				l = i
				lWall = true
			}
		} else {
			if h >= height[l] {
				l = i
				lWall = true
			}
		}

	}

	return trappedWater
}
