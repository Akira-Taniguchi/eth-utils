package eth_utils

import (
	//"fmt"
	//"math"
	//"reflect"
	//"strings"
	"errors"
	"github.com/shopspring/decimal"
)

const MinWei = "1"
// 2の256乗-1
const MaxWei = "115792089237316195423570985008687907853269984665640564039457584007913129639935"


func getEtherUnit(unit string) (decimal.Decimal, error) {
	var v decimal.Decimal
	var err error
	switch unit {
	case "wei":        v, err = decimal.NewFromString("1")
	case "kwei":       v, err = decimal.NewFromString("1000")
	case "babbage":    v, err = decimal.NewFromString("1000")
	case "femtoether": v, err = decimal.NewFromString("1000")
	case "mwei":       v, err = decimal.NewFromString("1000000")
	case "lovelace":   v, err = decimal.NewFromString("1000000")
	case "picoether":  v, err = decimal.NewFromString("1000000")
	case "gwei":       v, err = decimal.NewFromString("1000000000")
	case "shannon":    v, err = decimal.NewFromString("1000000000")
	case "nanoether":  v, err = decimal.NewFromString("1000000000")
	case "nano":       v, err = decimal.NewFromString("1000000000")
	case "szabo":      v, err = decimal.NewFromString("1000000000000")
	case "microether": v, err = decimal.NewFromString("1000000000000")
	case "micro":      v, err = decimal.NewFromString("1000000000000")
	case "finney":     v, err = decimal.NewFromString("1000000000000000")
	case "milliether": v, err = decimal.NewFromString("1000000000000000")
	case "milli":      v, err = decimal.NewFromString("1000000000000000")
	case "ether":      v, err = decimal.NewFromString("1000000000000000000")
	case "kether":     v, err = decimal.NewFromString("1000000000000000000000")
	case "grand":      v, err = decimal.NewFromString("1000000000000000000000")
	case "mether":     v, err = decimal.NewFromString("1000000000000000000000000")
	case "gether":     v, err = decimal.NewFromString("1000000000000000000000000000")
	case "tether":     v, err = decimal.NewFromString("1000000000000000000000000000000")
	default: {
		v, _ = decimal.NewFromString("0")
		err = errors.New("illegal unit string")
	}
	}
	return v, err
}


func FromWei(number decimal.Decimal, unit string) (decimal.Decimal, error) {
	v, err := getEtherUnit(unit)
	zero, _ := decimal.NewFromString("0")
	if err != nil {
		return zero, errors.New("illegal unit string")
	}
	if number.Equal(zero){
		return zero, nil
	}
	min, _ := decimal.NewFromString(MinWei)
	if number.Cmp(min) == -1{
		return zero, errors.New("illegal number")
	}
	max, _ := decimal.NewFromString(MaxWei)
	if number.Cmp(max) == 1{
		return zero, errors.New("illegal number")
	}
	return number.Div(v), nil
}
