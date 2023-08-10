package main

import (
	"github.com/akriventsev/go-card"
	"fmt"
)



func main(){

	card:=creditcard.Card{
		Type:        "Something",
        Number:      "5555555555554444 ",
        ExpiryMonth: 11,
        ExpiryYear:  2029,
        CVV:         "123",
	}
	validation:=card.Validate()
	if validation.Errors==nil{
		fmt.Println("This card is valid")
	} else {
		fmt.Println("This card is not valid")
	}
	// fmt.Printf("%+v\n", validation)
    // fmt.Printf("%+v\n", validation.Card)
}