package main

import (
	"fmt"
	"lectionmethods/pkg/card"
)

func main() {
	//svc := card.NewService("Netology Bank")
	//visa := card.IssueCard(svc, "Visa", "RUB")
	//master := svc.IssueCard("MasterCard", "USD")
	//
	//fmt.Println(visa)
	//fmt.Println(master)

	svc := card.NewService("Netology Bank")
	svcVal := *svc // разыменование указателя

	visa := (*svc).IssueCard("Visa", "RUB")
	master := svcVal.IssueCard("MasterCard", "USD")

	fmt.Println(visa.Name)
	fmt.Println(visa.Owner.Name)

	fmt.Println(visa)
	fmt.Println(master)
}
