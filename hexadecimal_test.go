package eth_utils

import (
	"testing"
)

func Test_Remove0xPrefix_1(t *testing.T) {
	v := Remove0xPrefix("0xe89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v != "e89943Ec20d856F064A87349a00Ef6aB00AED042" {
		t.Error("error")
	}
}

func Test_Remove0xPrefix_2(t *testing.T) {
	v := Remove0xPrefix("e89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v != "e89943Ec20d856F064A87349a00Ef6aB00AED042" {
		t.Error("error")
	}
}
