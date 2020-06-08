package main

import (
	"fmt"
	"lectionerrors/pkg/card"
)

func main() {
	svc := card.NewService("Netology Bank")

	//c := svc.SearchByNumber("0001")
	//fmt.Println(c.Balance)
	//fmt.Println("finish")

	//svc.Transfer(1, "0001", 100_00)
	//fmt.Println("finish")

	err := svc.Transfer(1, "0001", 100_00)
	if err != nil {
		switch err {
		case card.ErrSourceCardNotFound:
			fmt.Println("Sorry, can't complete transaction")
		case card.ErrTargetCardNotFound:
			fmt.Println("Please check target card number")
		default:
			fmt.Println("Something bad happened. Try again later")
		}
	}
}


