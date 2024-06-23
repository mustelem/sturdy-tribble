package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := [26]int{}

	for _, v := range magazine {
		order := v % 'a'
		letters[order]++
	}

	for _, v := range ransomNote {
		order := v % 'a'
		if letters[order] == 0 {
			return false
		}
		letters[order]--
	}

	return true
}
