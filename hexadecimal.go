package eth_utils

import (
	"regexp"
	"strings"
)

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

func isHex(value string) bool {
	if strings.ToLower(value) == "0x" {
		return true
	}
	removed := Remove0xPrefix(value)
	reg := regexp.MustCompile(`^[a-fA-F0-9]*$`)
	return reg.Match([]byte(removed))
}
