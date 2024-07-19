package main

func candy(ratings []int) int {
	var retVal int
	awards := make([]int, len(ratings))

	var size int = len(ratings)

	//Init
	for i := 0; i < size; i++ {
		awards[i] = 1
	}

	//Forward pass
	for i := 1; i < size; i++ {
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
