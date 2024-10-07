package main
import (
	"fmt"
)

func main() {
	fmt.Println(QuarterOfAYear(2))
	fmt.Println(QuarterOfAYear(5))
	fmt.Println(QuarterOfAYear(9))
	fmt.Println(QuarterOfAYear(11))
	fmt.Println(QuarterOfAYear(13))
	fmt.Println(QuarterOfAYear(-5))
}

func QuarterOfAYear(month int) int {
	fm:=0
	if month >= 2 && month<5{
		fm=1
	}else if month >= 5 && month<9{
		fm=(2)
	}else if month >= 9 && month<11{
		fm=(3)
	}else if month >= 11 && month<12 {
		fm=(4)
	}else {
		fm = -1
	}
	return fm
}