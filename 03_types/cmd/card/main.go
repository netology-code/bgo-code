package main

import (
	"fmt"
	"lectiontypes/pkg/card"
)

//func main() {
//	cardIssuer := "Visa"
//	cardBalance := int64(950_000_00)
//	cardCurrency := "RUB"
//	cardNumber := "0001"
//	cardIcon := "https://..."
//
//	println(cardIssuer, cardBalance, cardCurrency, cardNumber, cardIcon)
//}

//func main() {
//	card := struct {
//		id int64
//		issuer string
//		balance int64
//		currency string
//		number string
//		icon string
//	}{
//		id:       1,
//		issuer:   "MasterCard",
//		balance:  100_000_00,
//		currency: "RUB",
//		number:   "0001",
//		icon:     "https://...",
//	}
//
//	card.balance += 100_00
//	fmt.Println(card.balance)
//
//	card2 := struct {
//		id int64
//		issuer string
//		balance int64
//		currency string
//		number string
//		icon string
//	}{
//		id:       1,
//		issuer:   "MasterCard",
//		balance:  100_000_00,
//		currency: "RUB",
//		number:   "0001",
//		icon:     "https://...",
//	}
//
//	card = card2
//}

//func main() {
//	var c card.Card
//	fmt.Println(c)
//}

//func main() {
//	master := card.Card{
//		Id:       1,
//		Issuer:   "MasterCard",
//		Balance:  100_000_00,
//		Currency: "RUB",
//		Number:   "0001",
//		Icon:     "https://...",
//	}
//	card.Withdraw(master, 1000_00)
//	fmt.Println(master)
//}

//func main() {
//	master := card.Card{
//		Id:       1,
//		Issuer:   "MasterCard",
//		Balance:  100_000_00,
//		Currency: "RUB",
//		Number:   "0001",
//		Icon:     "https://...",
//	}
//	copy := master
//	copy.Balance -= 1000_00
//	fmt.Println(copy)
//	fmt.Println(master)
//}

//func main() {
//	master := card.Card{
//		Id:       1,
//		Issuer:   "MasterCard",
//		Balance:  100_000_00,
//		Currency: "RUB",
//		Number:   "0001",
//		Icon:     "https://...",
//	}
//	modified := card.Withdraw(master, 1000_00)
//	fmt.Println(modified)
//}

//func main() {
//	master := &card.Card{
//		Id:       1,
//		Issuer:   "MasterCard",
//		Balance:  100_000_00,
//		Currency: "RUB",
//		Number:   "0001",
//		Icon:     "https://...",
//	}
//	card.Withdraw(master, 1000_00)
//	fmt.Println(master)
//}

//func main() {
//	master := card.Card{
//		Id:       1,
//		Issuer:   "MasterCard",
//		Balance:  100_000_00,
//		Currency: "RUB",
//		Number:   "0001",
//		Icon:     "https://...",
//	}
//	visa := card.Card{
//		Id:       2,
//		Issuer:   "Visa",
//		Balance:  50_000_00,
//		Currency: "RUB",
//		Number:   "0002",
//		Icon:     "https://...",
//	}
//
//	//var cards [2]card.Card
//	//var cards [2]card.Card = [2]card.Card{master, visa}
//	//var cards [2]card.Card = [...]card.Card{master, visa}
//	//cards := [2]card.Card{}
//	//cards := [2]card.Card{master, visa}
//	cards := [...]card.Card{master, visa}
//
//	cards[0].Balance += 1000_00
//	fmt.Println(cards[0])
//
//	fmt.Println(len(cards))
//	fmt.Println(cards)
//}

func main() {
	master := card.Card{
		Id:       1,
		Issuer:   "MasterCard",
		Balance:  100_000_00,
		Currency: "RUB",
		Number:   "0001",
		Icon:     "https://...",
	}
	visa := card.Card{
		Id:       2,
		Issuer:   "Visa",
		Balance:  50_000_00,
		Currency: "RUB",
		Number:   "0002",
		Icon:     "https://...",
	}
	fmt.Println(master)
	fmt.Println(visa)

	//cards := []card.Card{master, visa}
	//cards := make([]card.Card, 5) // len = 5, cap = 5
	//cards := make([]card.Card, 5, 10) // len = 5, cap = 10
	//cards[0] = master
	//cards[1] = visa

	//fmt.Println(len(cards))
	//fmt.Println(cap(cards))
	//fmt.Println(cards[0])

	//cards := make([]card.Card, 0, 10) // len = 0, cap = 10
	//cards = append(cards, master)
	//cards = append(cards, visa)

	//fmt.Println(len(cards))
	//fmt.Println(cap(cards))
	//fmt.Println(cards)

	cards := []card.Card{master, visa}
	sum := card.Sum(cards)
	fmt.Println(sum)
}
