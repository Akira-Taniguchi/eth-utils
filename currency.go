package eth_utils

import (
	//"fmt"
	"math"
	//"reflect"
	"strings"
	"errors"
)

const MinWei = 0
// 2の256乗-1
const MaxWei = 115792089237316195423570985008687907853269984665640564039457584007913129639935

var ethUnit = map[string]float64 {
    "wei":          1,
    "kwei":         3,
    "babbage":      3,
    "femtoether":   3,
	"mwei":         6,
	"lovelace":     6,
	"picoether":    6,
	"gwei":         9,
	"shannon":      9,
	"nanoether":    9,
	"nano":         9,
	"szabo":        12,
	"microether":   12,
	"micro":        12,
	"finney":       15,
	"milliether":   15,
	"milli":        15,
	"ether":        18,
	"kether":       21,
	"grand":        21,
	"mether":       24,
	"gether":       27,
	"tether":       30,
}


func FromWei(number uint64, unit string) (float64, error) {
	v, exist := ethUnit[strings.ToLower(unit)]
	if exist == false {
		return 0, errors.New("illegal unit string")
	}
	// ほんまはMaxWeiとの比較も行いたいが、エラーになる。原因わからん。。。
	if number < MinWei{
		return 0, errors.New("illegal number")
	}
	//if number < MinWei || number > MaxWei{
	//	return 0, errors.New("illegal number")
	//}
	return float64(number) / math.Pow(10, v), nil
}
