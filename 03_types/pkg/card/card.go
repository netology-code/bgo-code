package card

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

//func Withdraw(card Card, amount int64) {
//	// TODO: checks
//	card.Balance -= amount
//}

//func Withdraw(card Card, amount int64) Card {
//	// TODO: checks
//	card.Balance -= amount
//	return card
//}

//func Withdraw(card *Card, amount int64) {
//	(*card).Balance -= amount
//}

func Withdraw(card *Card, amount int64) {
	card.Balance -= amount
}

func Sum(cards []Card) int64 {
	total := int64(0)
	for _, card := range cards {
		total += card.Balance
	}
	return total
}
