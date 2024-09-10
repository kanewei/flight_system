package calculator

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func MutipleString(multiplier string, multiplicand float64) (string, error) {

	multiplierFloat64, err := decimal.NewFromString(multiplier)
	if err != nil {
		return "", err
	}

	multiplicandFloat64 := decimal.NewFromFloat(multiplicand)

	return fmt.Sprintf("%v", multiplierFloat64.Mul(multiplicandFloat64).StringFixed(2)), nil
}
