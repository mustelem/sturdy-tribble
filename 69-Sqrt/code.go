package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	var lB, rB int

	lB = 1
	rB = x
	for lB < rB {
		prospective := lB + (rB-lB)/2
		squaredProspective := prospective * prospective
		switch {
		case squaredProspective == x:
			return prospective
		case squaredProspective > x:
			rB = prospective
		case squaredProspective < x:
			lB = prospective + 1
		}
	}

	return lB - 1
}
