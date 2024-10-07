package main

import (
	"fmt"
	"os"
)

func main() {
	are := os.Args[1]
	if len(os.Args) != 2 {
		fmt.Println()
	}
	result := ""
	for _, ele := range are {
		if ele >='a'&&ele<='z'{
			if ele == 'z' {
				result += string(ele-1) + string(ele)
			}else{
				result += string(ele) + string(ele+1)
			}
		}
		if ele>='A'&&ele<='Z'{
			if ele=='Z'{
				result+=string(ele-1)+string(ele)
			}else if ele=='A'{
				result+=string(ele)+string(ele+2)
			}else{
				result+=string(ele-1)+string(ele)+string(ele+1)
			}
		}
	}
	fmt.Println(result)
}
