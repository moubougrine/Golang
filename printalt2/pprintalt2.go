package main

import (
	"fmt"
)

func main() {
	for i := 'z'; i >= 'n'; i-- {
		fmt.Print(string(i))
		r:= i -'a'
		fmt.Print(string('z'-r))
	}
}
