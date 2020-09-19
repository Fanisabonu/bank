package card

import (
	"bank/pkg/bank/types"
)

//IssueCard -
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		Id:       1000,
		Pan:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

//Withdraw -
func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active != true {
		return card
	}
	if card.Balance < amount {
		return card
	}
	if amount < 0 {
		return card
	}
	if amount > 20_000_00 {
		return card
	}
	card.Balance = card.Balance - amount
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if card.Active == true {
		if amount >= 0 {
			if amount <= 50_00_000 {
				card.Balance = card.Balance + amount
			}
		}
	}
}

//AddBonus -
func AddBonus(card *types.Card, percent int, dayslnMonth int, dayslnYear int) {
	if card.Active != true {
		return
	}
	if card.Balance <= 0 {
		return
	}
	amount:=int(card.MinBalance)*percent/100*dayslnMonth/dayslnYear
	if amount>50_00_00{
		return 
	}
	card.Balance = card.Balance + types.Money(amount)

}

	func PaymentSources(cards []types.Card)(payments []types.PaymentSource) {
		for _, card := range cards {
			if card.Balance < 0 {
				continue
			}
			if !card.Active {
				continue
			}
			t:= types.PaymentSource{
				Type:"card",
				Number: string(card.Pan),
				Balance: card.Balance,
			}
			payments=append(payments,t)

		}
		return payments
	}
