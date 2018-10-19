package eth_utils

import (
	"fmt"
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
	fmt.Print("start")
	if strings.ToLower(value) == "0x" {
		return true
	}
	removed_0x := Remove0xPrefix(value)
	fmt.Print("start")
	for _, s := range removed_0x {
		fmt.Printf("name: %s\n", s)
	}
	return true
}
