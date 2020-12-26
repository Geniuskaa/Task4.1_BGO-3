package card

import "github.com/Geniuskaa/Task4.1_BGO-3/pkg/transaction"

type Card struct {
	Id int64
	Issuer string
	Currency string
	Balance int64
	Number string
	Transactions []*transaction.Transaction
}

type Service struct {
	bank string
	StoreOfCards []*Card
}

func NewService(storeOfCards []*Card, bankName string) *Service {
	return &Service{
		bank: bankName,
		StoreOfCards: storeOfCards}
}

func (s *Service) AddCard(id int64, issuer string, currency string, balance int64, number string) {
	s.StoreOfCards = append(s.StoreOfCards, &Card{
		Id:       id,
		Issuer:   issuer,
		Currency: currency,
		Balance:  balance,
		Number:   number,
	})
}

func (s *Service) SearchCards(number string) (available bool, index int) {
	for i, _ := range s.StoreOfCards {
		if s.StoreOfCards[i].Number == number {
			return true, i
		}
	}
	return false, -1
}


