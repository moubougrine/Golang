package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	are := os.Args[1]
	are1 := os.Args[2]
	str,_:=strconv.Atoi(are1)
	Atio,_:=strconv.Atoi(are)
	for i := str; i >= Atio; i-- {
		fmt.Print(i, " ")
	}
}
