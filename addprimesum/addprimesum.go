package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	are := os.Args[1]
	Atio,_ := strconv.Atoi(are)
	count := 0
	for i := 2; i <= Atio; i++ {
		if fPrime(i){
			count += i
		}
	}
	fmt.Println(count )

}
func fPrime(s int) bool {
	if s<=1{
		return false
	}
	for i := 2; i*i<= s; i++ {
		if s%i== 0 {
			return false
		}
	}
	return true
}
