package card

import (
	"errors"
	"github.com/Geniuskaa/Task4.1_BGO-3/pkg/transaction"
	"math/rand"
	"strings"
)

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

var ErrCardNotInOurBase = errors.New("Данной карты нет в нашей базе данных.")



func (s *Service) SearchCards(number string) (err error, index int) {
	for i, _ := range s.StoreOfCards {
		if s.StoreOfCards[i].Number == number {
			return nil, i
		}
	}
	if strings.HasPrefix(number, "5106 21") {
		s.StoreOfCards = append(s.StoreOfCards, &Card{
			Id:           rand.Int63n(1000),
			Issuer:       "VISA",
			Currency:     "RUB",
			Balance:      rand.Int63n(10000000),
			Number:       number,
			Transactions: nil,
		})
		return nil, len(s.StoreOfCards) - 1
	}
	return ErrCardNotInOurBase, -1
}


