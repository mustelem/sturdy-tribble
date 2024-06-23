package main

func canConstruct(ransomNote string, magazine string) bool {
	magArchive := map[rune]int{}

	for _, v := range magazine {
		count, e := magArchive[v]
		if e {
			magArchive[v] = count + 1
		} else {
			magArchive[v] = 1
		}
	}

	for _, v := range ransomNote {
		count, e := magArchive[v]

		if e {
			if count == 0 {
				return false
			} else {
				magArchive[v] = count - 1
			}
		} else {
			return false
		}
	}

	return true
}
