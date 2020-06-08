package card

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(issuer string, currency string) *Card {
	card := &Card{
		Issuer:   issuer,
		Balance:  0,
		Currency: currency,
		Number:   "0000 0000 0000 0000",
		Icon:     "https://...",
		Pin:      "0000",
	}
	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

func (s *Service) Sum() int64 {
	total := int64(0)
	for _, card := range s.Cards {
		total += card.Balance
	}
	return total
}

type Card struct {
	Owner
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
	Pin      string
}

type Owner struct {
	Name string
}

func IssueCard(service Service, issuer string, currency string) *Card {
	// TODO: check input
	card := &Card{
		Issuer:   issuer,
		Balance:  0,
		Currency: currency,
		Number:   "0001",
		Icon:     "https://...",
		Pin:      "0000",
	}
	service.Cards = append(service.Cards, card)
	return card
}

func Sum(cards []Card) int64 {
	total := int64(0)
	for _, card := range cards {
		total += card.Balance
	}
	return total
}

//func (m time.Month) Days() int {
//	// в качестве получателя нельзя использовать тип из другого пакета
//}
