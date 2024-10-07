package main

import "fmt"

func main() {
    fmt.Println(EvenLength([]string{"Hi", "Hola", "Ola", "Ciao", "Salut", "Hallo"}))
}
func EvenLength(strings []string) []string {
	result := []string{}
	for _,ele := range strings{
		if len(ele)%2==0{
			result = append(result, ele)
		}
	}
	return result
}
