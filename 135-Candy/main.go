package main

func candy(ratings []int) int {
	var retVal int
	var backtick int

	awards := make([]int, len(ratings))
	awards[0]++

	for i := 1; i < len(ratings); i++ {
		awards[i]++
		if ratings[i] == ratings[i-1] {
			continue
		}
		if ratings[i] < ratings[i-1] {
			if i == len(ratings)-1 {
				j := 0
				for j <= backtick {
					if ratings[i-j] < ratings[i-j-1] {
						awards[i-j-1] = awards[i-j] + 1
					}
					j++
				}
				backtick = 0
			} else {
				backtick++
			}

		} else {
			j := 0
			for j < backtick {
				if ratings[i-j-1] < ratings[i-j-2] {
					awards[i-j-2] = awards[i-j-1] + 1
				}
				j++
			}
			backtick = 0
			awards[i] = awards[i-1] + 1
		}
	}

	for _, v := range awards {
		retVal += v
	}

	return retVal
}
