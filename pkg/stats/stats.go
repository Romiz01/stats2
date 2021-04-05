package stats

import (
	"github.com/Romiz01/bank1/v2/pkg/types"
)

// Avg 1
func Avg(payments []types.Payment) types.Money {

	PaySums := types.Money(0)
	indexPay := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		PaySums += moneyPayments
		indexPay++
	}
	return PaySums / indexPay
}

// TotalInCategory 2
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	PaySum := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}

		PayMoney := payment.Amount
		PaySum += PayMoney

	}
	return PaySum
}
