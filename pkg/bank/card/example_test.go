package card

import (
	"fmt"
	"bank/pkg/bank/types"
)
func ExampleWithdraw_positive(){
	result := Withdraw(types.Card{Balance:10000, Active: true}, -100)
	fmt.Println (result.Balance)

	result = Withdraw(types.Card{Balance:10000,Active: false}, 100)
	fmt.Println (result.Balance)
	
	result = Withdraw(types.Card{Balance:10000, Active: true}, 10001)
	fmt.Println (result.Balance)
	
	result = Withdraw(types.Card{Balance:40_000_00, Active: true}, 20_000_01)
	fmt.Println (result.Balance)

	result = Withdraw(types.Card{Balance:40_000_00, Active: true}, 10_000_00)
	fmt.Println (result.Balance)

	//Output: 
	//10000
	//10000
	//10000
	//4000000
	//3000000
}

func ExampleDeposit_positive() {
	card := types.Card{
	  Active: true,
	  Balance: 20_000_00,
	}
  
	Deposit(&card, -10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
  }
  func ExampleDeposit_nowActive() {
	card := types.Card{
	  Active: false,
	  Balance: 20_000_00,
	}
  
	Deposit(&card, -10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
  }
  func ExampleDeposit_limit() {
	card := types.Card{
	  Active: true,
	  Balance: 20_000_00,
	}
  
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)
	// Output: 2000000
  }

  func ExamplePaymentSources()  {
	cards:=[]types.Card{
		{
			Pan:"5058",
			Active: true,
			Balance:50,
		},
		{
			Pan:"5058",
			Active: false,
			Balance:50,
		},
		{
			Pan:"5058",
			Active: true,
			Balance:-50,
		},
	}
	payments:=PaymentSources(cards)
for _, payment := range payments {
	
	fmt.Println(payment.Number)
}
//Output:
//5058
}