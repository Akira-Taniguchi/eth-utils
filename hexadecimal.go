package eth_utils

import "strings"

func Remove0xPrefix(value string) string {
	if strings.HasPrefix(strings.ToUpper(value), "0X") == false {
		return value
	}
	return string([]rune(value)[2:])
}
