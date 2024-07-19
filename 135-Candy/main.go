package main

func candy(ratings []int) int {
	var retVal int
	awards := make([]int, len(ratings))

	var size int = len(ratings)

	//Init
	awards[0] = 1

	//Forward pass
	for i := 1; i < size; i++ {
		awards[i] = 1
		if ratings[i] > ratings[i-1] {
			awards[i] = awards[i-1] + 1
		}
	}

	//Backward pass
	for i := size - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			awards[i-1] = max(awards[i-1], awards[i]+1)
		}
	}

	for _, v := range awards {
		retVal += v
	}

	return retVal
}
