package card

import "errors"

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

func (s *Service) FindByNumber(number string) (*Card, bool) {
	for _, c := range s.Cards {
		if c.Number == number {
			return c, true
		}
	}
	return nil, false
}

func (s *Service) SearchById(id int64) *Card {
	for _, card := range s.Cards {
		if card.Id == id {
			return card
		}
	}
	return nil
}

func (s *Service) FindById(id int64) (*Card, bool) {
	for _, c := range s.Cards {
		if c.Id == id {
			return c, true
		}
	}
	return nil, false
}

//func (s *Service) Transfer(fromId int64, toNumber string, amount int64) {
//	source := s.SearchById(fromId)
//	target := s.SearchByNumber(toNumber)
//
//	source.Balance -= amount
//	target.Balance += amount
//}

//func (s *Service) Transfer(fromId int64, toNumber string, amount int64) {
//	if source, ok := s.FindById(fromId); ok {
//		if target, ok := s.FindByNumber(toNumber); ok {
//			source.Balance -= amount
//			target.Balance += amount
//		}
//	}
//}

//type TransferError string
//
//func (e *TransferError) Error() string {
//	return string(*e)
//}
//
//func (s *Service) Transfer(fromId int64, toNumber string, amount int64) error {
//	source, ok := s.FindById(fromId)
//	if !ok {
//		err := TransferError("source card not found")
//		return &err
//	}
//	target, ok := s.FindByNumber(toNumber)
//	if !ok {
//		err := TransferError("target card not found")
//		return &err
//	}
//	// в этой точке уверены, что все проверки пройдены
//	source.Balance -= amount
//	target.Balance += amount
//	return nil
//}

//type TransferError struct {
//	message string
//}
//
//func (e *TransferError) Error() string {
//	return e.message
//}
//
//func (s *Service) Transfer(fromId int64, toNumber string, amount int64) error {
//	source, ok := s.FindById(fromId)
//	if !ok {
//		return &TransferError{"source card not found"}
//	}
//	target, ok := s.FindByNumber(toNumber)
//	if !ok {
//		return &TransferError{"target card not found"}
//	}
//	// в этой точке уверены, что все проверки пройдены
//	source.Balance -= amount
//	target.Balance += amount
//	return nil
//}

var (
	ErrSourceCardNotFound = errors.New("source card not found")
	ErrTargetCardNotFound = errors.New("target card not found")
)

func (s *Service) Transfer(fromId int64, toNumber string, amount int64) error {
	source, ok := s.FindById(fromId)
	if !ok {
		return ErrSourceCardNotFound
	}
	target, ok := s.FindByNumber(toNumber)
	if !ok {
		return ErrTargetCardNotFound
	}
	// в этой точке уверены, что все проверки пройдены
	source.Balance -= amount
	target.Balance += amount
	return nil
}

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
	Pin      string
}
