package transfer

import (
	"errors"
	"fmt"
	"github.com/Geniuskaa/Task4.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task4.1_BGO-3/pkg/transaction"
	"math/rand"
	"time"
)

type Service struct {  // Вроде понятно, что константы тут использовать неуместно, но как правильно можно было бы
	CardSvc           *card.Service // оформить эту структуру, чтобы и константы и изменяемые поля уживались вместе, а то я
	toTinkPercent     float64       // попытался использовать константы, но не получилось(
	fromTinkPercent   float64
	fromTinkMinSum    int64
	otherCardsPercent float64
	otherCardsMinSum  int64
}

func NewService(cardSvc *card.Service, toTinPer float64, fromTinPer float64, fromTinMSum int64, otherCardPer float64, otherCardMSum int64) *Service {
	return &Service{
		CardSvc:           cardSvc,
		toTinkPercent:     toTinPer,
		fromTinkPercent:   fromTinPer / 100,
		fromTinkMinSum:    fromTinMSum,
		otherCardsPercent: otherCardPer / 100,
		otherCardsMinSum:  otherCardMSum,
	}
}

var (
	ErrMoneyOnCardOfSenderDontEnough = errors.New("На карте отправителя баланс меньше суммы перевода.")
	ErrTooLowSumOfTransfer = errors.New("Слишком маленькая сумма перевода.")
)

func (s *Service) Card2Card(from, to string, amount int64) (total int64, err error) {
	amountInCents := amount * 100
	errOfFrom, indexOfFrom := s.CardSvc.SearchCards(from)
	errOfTo, indexOfTo := s.CardSvc.SearchCards(to)
	if errOfFrom != nil {
		fmt.Println("Карты с которой вы хотите выполнить перевод нет в нашей базе данных.")
		return 0, card.ErrCardNotInOurBase
	}
	if errOfTo != nil {
		fmt.Println("Карты с которой вы хотите выполнить перевод нет в нашей базе данных.")
		return 0, card.ErrCardNotInOurBase
	}
	// При новых двух условиях, которые я написал около половины кода этой функции просто стало мусором.
	// Получается обработка карт "чужих" банков нам не нужна?

	if errOfFrom == nil {
		if s.CardSvc.StoreOfCards[indexOfFrom].Balance > amountInCents { // Проверяем хватает ли денег на балансе
			if errOfTo == nil {
				s.addTransaction(indexOfFrom, amount)
				s.CardSvc.StoreOfCards[indexOfFrom].Balance -= amountInCents
				s.CardSvc.StoreOfCards[indexOfTo].Balance += amountInCents
				return amount, nil
			} else {
				if amountInCents > s.fromTinkMinSum { // Проверяем больше ли сумма перевода чем минимальная по тарифу
					total := int64(float64(amountInCents) * (1.0 + s.fromTinkPercent / 100))
					s.addTransaction(indexOfFrom, amount)
					s.CardSvc.StoreOfCards[indexOfFrom].Balance -= total
					return total / 100, nil
				} else {
					fmt.Println("Слишком маленькая сумма перевода, введите сумму более 10 руб!")
					return 0, ErrTooLowSumOfTransfer
				}
			}
		}
		fmt.Println("Недостаточно средств на балансе вашей карты.")
		return 0, ErrMoneyOnCardOfSenderDontEnough
	}

	if errOfTo == nil {
		s.CardSvc.StoreOfCards[indexOfTo].Balance += amountInCents
		return amount, nil
	}

	if amountInCents > s.otherCardsMinSum {
		total := int64(float64(amountInCents) * (1 + s.otherCardsPercent / 100))
		return total / 100, nil
	}

	fmt.Println("Сумма перевода меньше минимального значения! Перевод невозможен.")
	return 0, ErrTooLowSumOfTransfer
}

func (s *Service) addTransaction(index int, amount int64) {
		s.CardSvc.StoreOfCards[index].Transactions = append(s.CardSvc.StoreOfCards[index].Transactions, &transaction.Transaction{
		Id:     rand.Int63n(20),
		Amount: amount,
		MCC:    "5090",
		Date:   time.Now().Unix(),
		Status: "Completed",
	})

}

