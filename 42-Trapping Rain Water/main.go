package main

func trap(height []int) int {
	l, r, trappedWater := 0, 0, 0
	var temp []int = make([]int, len(height))

	rWall := false

	for i := 0; i < len(height); i++ {

		// Update right
		if i > 0 && height[i-1] <= height[i] {
			r = i
			rWall = true
		}

		// Calculate possible accumulation
		if height[i] < height[l] {
			temp[i] = height[l] - height[i]
		}

		// Calculate actual accumulation
		if rWall && r-l > 1 {
			minH := min(height[l], height[r])

			for ptr := l + 1; ptr < r; ptr++ {
				if height[ptr] >= minH {
					continue
				}
				canHoldAdditionally := min(minH-height[ptr], temp[ptr])
				trappedWater += canHoldAdditionally
				temp[ptr] -= canHoldAdditionally
				height[ptr] += canHoldAdditionally
			}
		}

		// Update left
		if i == r && height[l] <= height[i] {
			l = i
			rWall = false
		}

	}

	return trappedWater
}
