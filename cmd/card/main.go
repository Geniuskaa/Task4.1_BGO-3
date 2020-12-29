package main

import (
	"fmt"
	"github.com/Geniuskaa/Task4.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task4.1_BGO-3/pkg/transfer"
)

func main() {
	bank := card.NewService([]*card.Card{},"Tinkoff")
	bank.AddCard(1,"VISA", "RUB", 14_800_00, "4724 3728 3929 5030")
	bank.AddCard(2, "MASTER", "RUB", 28_750_00, "6930 2857 3892 2967")
	bank.AddCard(3, "VISA", "RUB", 352_362_00, "4626 9205 2859 2852")

	transfers := transfer.NewService(bank, 0, 0.5, 10_00, 1.5, 30_00)
	_, err := transfers.Card2Card("4724 3728 3929 5030", "6930 2857 3892 2967", 5_425)
	if err != nil {
		switch err {
		case transfer.ErrMoneyOnCardOfSenderDontEnough:
			fmt.Println("Недостаточно средств на балансе для перевода.")
		case transfer.ErrTooLowSumOfTransfer:
			fmt.Println("Слишком маленькая сумма перевода.")
		default:
			fmt.Println("Возникла непредвиденная ошибка.")
		}
	}
	//transfers.Card2Card("4724 3708 3929 5030", "6930 2857 3892 2967", 50_425)
	//transfers.Card2Card("4724 3728 3929 5030", "6930 2857 3812 2967", 725)
	//transfers.Card2Card("2424 3728 2829 5030", "97030 2857 3892 2967", 38_425)





}
