package main

func strStr(haystack string, needle string) int {
	var hayIdx int = -1

	if len(haystack) < len(needle) {
		return hayIdx
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

	return hayIdx
}
