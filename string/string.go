package main

import "fmt"

func main(){
	var str string
	fmt.Scanln(&str)
	sampleRune:= []rune(str)
	flag:=0
	for i:=0;i<len(str);i+=2 {
		if flag==0{
			fmt.Printf("%c%c",sampleRune[i],sampleRune[i+1])
			flag=1
		} else if flag==1 {
			fmt.Printf("%c%c",sampleRune[i+1],sampleRune[i])
			flag=0
		}
	}
	fmt.Printf("\n")
	
	
}