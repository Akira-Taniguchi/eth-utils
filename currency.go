package eth_utils

import (
	"errors"
	"math"
	"reflect"
)

const MinWei = 0
// 2の256乗-1
const MaxWei = 115792089237316195423570985008687907853269984665640564039457584007913129639935

var ethUnit = map[string]uint64 {
    "wei":          1,
    "kwei":         1000,
    "babbage":      1000,
    "femtoether":   1000,
	"mwei":         1000000,
	"lovelace":     1000000,
	"picoether":    1000000,
	"gwei":         1000000000,
	"shannon":      1000000000,
	"nanoether":    1000000000,
	"nano":         1000000000,
	"szabo":        1000000000000,
	"microether":   1000000000000,
	"micro":        1000000000000,
	"finney":       1000000000000000,
	"milliether":   1000000000000000,
	"milli":        1000000000000000,
	"ether":        1000000000000000000,
	"kether":       1000000000000000000000,
	"grand":        1000000000000000000000,
	"mether":       1000000000000000000000000,
	"gether":       1000000000000000000000000000,
	"tether":       1000000000000000000000000000000,
}


func FromWei(number uint64, unit string) (float64, error) {
	ethUnitValue, ok = ethUnit[unit]
	if ok == false {
		keys := reflect.ValueOf(ethUnit).MapKeys()
		return 0, errors.New(fmt.Printf("Unknown unit.  Must be one of %s", keys))
	}
	if number == 0 {
		return 0, nil
	}
	if number < MinWei || number > MaxWei{
		return 0, errors.New(fmt.Printf("value must be between 1 and %s", MaxWei))
	}
	return uint64 / ethUnitValue, nil
}
