package main

import (
	"fmt"
	"os"
	"strconv"
)

func intToRoman(num int) string {
	var retVal string
	var ptr int

	for num > 0 {
		number := strconv.Itoa(num)
		sign, err := strconv.Atoi(string(number[ptr]))
		if err != nil {
			fmt.Printf("Error encountered: %s\n", err)
			os.Exit(1)
		}
		pow := len(number) - ptr

		switch sign {
		case 9:
			switch pow {
			case 3:
				retVal += "CM"
				num -= 900
			case 2:
				retVal += "XC"
				num -= 90
			case 1:
				retVal += "IX"
				num -= 9
			}
		case 8, 7, 6, 5:
			switch pow {
			case 3:
				retVal += "D"
				num -= 500
			case 2:
				retVal += "L"
				num -= 50
			case 1:
				retVal += "V"
				num -= 5
			}
		case 4:
			switch pow {
			case 3:
				retVal += "CD"
				num -= 400
			case 2:
				retVal += "XL"
				num -= 40
			case 1:
				retVal += "IV"
				num -= 4
			}
		case 3, 2, 1:
			switch pow {
			case 4:
				retVal += "M"
				num -= 1000
			case 3:
				retVal += "C"
				num -= 100
			case 2:
				retVal += "X"
				num -= 10
			case 1:
				retVal += "I"
				num -= 1
			}

		}

	}

	return retVal
}
