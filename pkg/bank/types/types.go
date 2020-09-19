package types

type Money int64
type Currency string 

const(
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Pan string 

type Card struct {
	Id	int
	Pan	Pan
	Balance Money
	MinBalance Money 
	Currency Currency
	Color string
	Name string
	Active bool 
}
type PaymentSource struct{
	Type string
	Number string 
	Balance Money 
		
}