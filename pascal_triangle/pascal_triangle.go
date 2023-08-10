package main

import "fmt"

func pascal(input int){
	for line:=1;line<=input;line++ {
		for space:=1;space<=input-line;space++ {
			fmt.Printf(" ")
		}
		coef:=1
		for i:=1;i<=line;i++ {
			fmt.Printf("%d ",coef)
			coef=coef*(line-i)/i
		}
		fmt.Printf("\n")
	}
}
func main(){
	var input int
	fmt.Scanln(&input)
	pascal(input)
	

}