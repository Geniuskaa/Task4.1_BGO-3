package transfer

import (
	"github.com/Geniuskaa/Task4.1_BGO-3/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc           *card.Service
		toTinkPercent     float64
		fromTinkPercent   float64
		fromTinkMinSum    int64
		otherCardsPercent float64
		otherCardsMinSum  int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}

	bank := card.NewService([]*card.Card{}, "TestBank")
	bank.StoreOfCards = append(bank.StoreOfCards, &card.Card{
		Id:           1,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      103_373_93,
		Number:       "3920 4923 3976 3972",
		Transactions: nil,
	})
	bank.StoreOfCards = append(bank.StoreOfCards, &card.Card{
		Id:           2,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      10_650_00,
		Number:       "9520 4923 3914 3871",
		Transactions: nil,
	})
	bank.StoreOfCards = append(bank.StoreOfCards, &card.Card{
		Id:           3,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      1_030_20,
		Number:       "2920 0723 3976 1272",
		Transactions: nil,
	})
	slice := fields{
		CardSvc:           bank,
		toTinkPercent:     0,
		fromTinkPercent:   0.5,
		fromTinkMinSum:    10_00,
		otherCardsPercent: 1.5,
		otherCardsMinSum:  30_00,
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{"From us to us, enough money", slice, args{
			from:   "3920 4923 3976 3972",
			to:     "9520 4923 3914 3871",
			amount: 15_464,
		}, 15_464, true},
		{"From us to us, no money", slice, args{
			from:   "2920 0723 3976 1272",
			to:     "9520 4923 3914 3871",
			amount: 146_346,
		}, 0, false},
		{"From us to other, enough money", slice, args{
			from:   "9520 4923 3914 3871",
			to:     "3952 2819 1289 2463",
			amount: 5_346,
		}, 5_372, true},
		{"From us to other, no money", slice, args{
			from:   "2920 0723 3976 1272",
			to:     "9392 0270 9723 3902",
			amount: 9_482,
		}, 0, false},
		{"From other to us", slice, args{
			from:   "4380 2086 0829 2071",
			to:     "3920 4923 3976 3972",
			amount: 293_387,
		}, 293_387, true},
		{"From other to other", slice, args{
			from:   "3860 9743 2983 8638",
			to:     "9037 2387 2990 1974",
			amount: 10,
		}, 0, false},
	}
	for _, tt := range tests {
			s := &Service{
				CardSvc:           tt.fields.CardSvc,
				toTinkPercent:     tt.fields.toTinkPercent,
				fromTinkPercent:   tt.fields.fromTinkPercent,
				fromTinkMinSum:    tt.fields.fromTinkMinSum,
				otherCardsPercent: tt.fields.otherCardsPercent,
				otherCardsMinSum:  tt.fields.otherCardsMinSum,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
	}
}