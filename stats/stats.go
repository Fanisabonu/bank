package stats
import (
	"github.com/Fanisabonu/bank/pkg/bank/types"
)

//Avg -
func Avg(payments []types.Payment)(money types.Money) {
	for _, payment := range payments {
		money += payment.Amount
	}
	return money / types.Money(len(payments))
}
