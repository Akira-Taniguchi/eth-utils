package eth_utils

import (
//"strconv"
)

func IsHexAddress(address string) bool {
	// i := strconv.FormatInt(address, 16)
	// strconv.ParseInt(i, 16, 32)
	//あたまの0xけして、40文字かチェックする
	return true
}

//def is_hex_address(value: Any) -> bool:
//"""
//Checks if the given string of text type is an address in hexadecimal encoded form.
//"""
//if not is_text(value):
//return False
//elif not is_hex(value):
//return False
//else:
//unprefixed = remove_0x_prefix(value)
//return len(unprefixed) == 40
