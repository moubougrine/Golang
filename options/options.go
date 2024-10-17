package main 

import (
	"os"
	"fmt"
)

func  main()  {
	char:="******zy xwvutsrq ponmlkji hgfedcba"
	char1:=[]string{"0","0","0","0","0","0","0","0"," ","0","0","0","0","0","0","0","0"," ","0","0","0","0","0","0","0","0"," ","0","0","0","0","0","0","0","0"}
	
	if len(os.Args)>3{
		return
	}
	
	if len(os.Args) == 2 {
		a1:=os.Args[1]
		for i,ele:= range a1{
			
			if len(a1)==1 && (i== 1 && !IfAlphabe(ele)){
				fmt.Println("Invalid Option")
				return
			}
			
			if i == 0 && ele=='-'{
				
				for i,ele1:= range a1[1:]{
					
					if i == 0 && ele1=='h' {
						fmt.Print("options: abcdefghijklmnopqrstuvwxyz")
						return
					}else{
						
						for j,ele2:=range char{
							if ele1==ele2{
								char1[j]="1"
							}
						}
					}
				}
			}
		}
	}
	if len(os.Args)==3 {
		a1:=os.Args[1]
		a2:=os.Args[2]
		
		for i,ele:= range a1{
			
			if len(a1)==1 && (i== 1 && !IfAlphabe(ele)){
				fmt.Println("Invalid Option")
				return
			}
			
			if i == 0 && ele=='-'{
				for i,ele1:= range a1[1:]{
					
					if i == 0 && ele1=='h' {
						fmt.Print("options: abcdefghijklmnopqrstuvwxyz")
						return
					}else{
						
						for j,ele2:=range char{
							if ele1==ele2{
								char1[j]="1"
							}
						}
					}
				}
			}
		}
		for i,ele:= range a2{
			
			if len(a2)==1 && (i== 1 && !IfAlphabe(ele)){
				fmt.Println("Invalid Option")
				return
			}
			
			if i == 0 && ele=='-'{
				for i,ele1:= range a2[1:]{
					
					if i == 0 && ele1=='h' {
						fmt.Print("options: abcdefghijklmnopqrstuvwxyz")
						return
					}else{
						
						for j,ele2:=range char{
							if ele1==ele2{
								char1[j]="1"
							}
						}
					}
				}
			}
		}
	}
	for _,elle:=range char1{
		fmt.Print(string(elle))
	}
}

func IfAlphabe(a rune) bool{
	if a >= 'a' && a <= 'z' {
		return true
	}
	return false
}