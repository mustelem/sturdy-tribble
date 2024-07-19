package main

func candy(ratings []int) int {
	var retVal int
	var backtrack int

	awards := make([]int, len(ratings))
	awards[0]++

	for i := 1; i < len(ratings); i++ {
		awards[i]++
		if ratings[i] == ratings[i-1] {
			if backtrack > 0 {
				j := 0
				for j < backtrack {
					if ratings[i-j-1] < ratings[i-j-2] {
						if i-j-3 >= 0 && ratings[i-j-2] > ratings[i-j-3] {
							awards[i-j-2] = max(awards[i-j-1]+1, awards[i-j-3]+1)
						} else {
							awards[i-j-2] = awards[i-j-1] + 1
						}
					}
					j++
				}
				backtrack = 0
			}
			continue
		}
		if ratings[i] < ratings[i-1] {
			if i == len(ratings)-1 {
				j := 0
				for j <= backtrack {
					if ratings[i-j] < ratings[i-j-1] {
						if i-j-2 >= 0 && ratings[i-j-1] > ratings[i-j-2] {
							awards[i-j-1] = max(awards[i-j]+1, awards[i-j-2]+1)
						} else {
							awards[i-j-1] = awards[i-j] + 1
						}
					}
					j++
				}
				backtrack = 0
			} else {
				backtrack++
			}

		} else {
			j := 0
			for j < backtrack {
				if ratings[i-j-1] < ratings[i-j-2] {
					if i-j-3 >= 0 && ratings[i-j-2] > ratings[i-j-3] {
						awards[i-j-2] = max(awards[i-j-1]+1, awards[i-j-3]+1)
					} else {
						awards[i-j-2] = awards[i-j-1] + 1
					}
				}
				j++
			}
			backtrack = 0
			awards[i] = awards[i-1] + 1
		}
	}

	for _, v := range awards {
		retVal += v
	}

	return retVal
}
