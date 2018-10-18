package eth_utils

import "strings"

func Remove0xPrefix(value string) string {
	if Is0xPrefixed(value) == false {
		return value
	}
	return string([]rune(value)[2:])
}

func Is0xPrefixed(value string) bool {
	return strings.HasPrefix(strings.ToLower(value), "0x")
}

func Add0xPrefix(value string) string {
	if Is0xPrefixed(value) {
		return value
	}
	return "0x" + value
}
