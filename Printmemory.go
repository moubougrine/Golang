package piscine

import "fmt"

func PrintMemory(arr [10]byte) {
	ToHix := func(str int) string {
		if str == 0 {
			return "00"
		}
		res := ""
		char := "0123456789abcdef"
		for str > 0 {
			res = string(char[str%16]) + res
			str /= 16
		}
		return res
	}
	result := ""
	res := []string{}
	for _, ele := range arr {
		res = append(res, ToHix(int(ele)))
		if ele >= 32 {
			result += string(ele)
		} else {
			result += "."
		}
	}
	for i := range res {
		if (i+1)%4 == 0 {
			fmt.Println(res[i])
		} else {
			if i != len(res)-1 {
				fmt.Print(res[i], " ")
			} else {
				fmt.Print(res[i])
			}
		}
	}
	fmt.Println()
	fmt.Print(result)
}
