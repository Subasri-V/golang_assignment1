package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string
	fmt.Scanln(&str)
	newString:=[]rune(str)
	substr:="go"
	ctr:=strings.Count(string(newString),substr)
	
	fmt.Println(ctr)
}