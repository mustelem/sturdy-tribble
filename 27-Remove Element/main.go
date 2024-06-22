package main

func main() {

}

func removeElement(nums []int, val int) int {
	spaces := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			spaces = append(spaces, i)
		} else {
			if i > 0 && len(spaces) > 0 {
				ptr := spaces[0]
				spaces = spaces[1:]
				nums[ptr] = nums[i]
				spaces = append(spaces, i)
			}
		}
	}

	return len(nums) - len(spaces)
}
