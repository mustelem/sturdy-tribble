package main

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	needlePtr := 0
	for hayPtr := 0; hayPtr <= len(haystack)-len(needle); hayPtr++ {
		if needle[needlePtr] == haystack[hayPtr+needlePtr] {
			for needlePtr < len(needle) && needle[needlePtr] == haystack[hayPtr+needlePtr] {
				needlePtr++
			}
			if needlePtr == len(needle) {
				return hayPtr
			} else {
				needlePtr = 0
			}
		} else {
			needlePtr = 0
		}
	}

	return -1
}
