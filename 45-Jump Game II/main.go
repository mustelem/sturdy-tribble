package main

func jump(nums []int) int {
	closests := make([]int, 0, 100)
	closests = append(closests, len(nums)-1)

	for i := len(nums) - 2; i >= 0; i-- {
		for rIn := 0; rIn < len(closests); rIn++ {
			if i+nums[i] >= closests[rIn] {
				// Reaches target
				steps := rIn + 1
				if steps < len(closests) {
					closests[steps] = i
				} else {
					closests = append(closests, i)
				}
				if i == 0 {
					return steps
				}
				break
			}
		}

	}

	return 0
}
