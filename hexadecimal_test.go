package eth_utils

import (
	"testing"
)

func TestRemove0xPrefix1(t *testing.T) {
	v := Remove0xPrefix("0xe89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v != "e89943Ec20d856F064A87349a00Ef6aB00AED042" {
		t.Error("error")
	}
}

func TestRemove0xPrefix2(t *testing.T) {
	v := Remove0xPrefix("e89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v != "e89943Ec20d856F064A87349a00Ef6aB00AED042" {
		t.Error("error")
	}
}

func TestIs0xPrefixed1(t *testing.T) {
	v := Is0xPrefixed("e89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v {
		t.Error("error")
	}
}

func TestIs0xPrefixed2(t *testing.T) {
	v := Is0xPrefixed("0xe89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v == false {
		t.Error("error")
	}
}

func TestIs0xPrefixed3(t *testing.T) {
	v := Is0xPrefixed("0Xe89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v == false {
		t.Error("error")
	}
}

func TestAdd0xPrefix1(t *testing.T) {
	v := Add0xPrefix("e89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v != "0xe89943Ec20d856F064A87349a00Ef6aB00AED042" {
		t.Error("error")
	}
}

func TestAdd0xPrefix2(t *testing.T) {
	v := Add0xPrefix("0xe89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v != "0xe89943Ec20d856F064A87349a00Ef6aB00AED042" {
		t.Error("error")
	}
}

func TestIsHex1(t *testing.T) {
	v := isHex("0xe89943Ec20d856F064A87349a00Ef6aB00AED042")
	if v == false {
		t.Error("error")
	}
}

func TestIsHex2(t *testing.T) {
	v := isHex("0x")
	if v == false {
		t.Error("error")
	}
}

func TestIsHex3(t *testing.T) {
	v := isHex("0xZ")
	if v {
		t.Error("error")
	}
}

func TestIsHex4(t *testing.T) {
	v := isHex("123")
	if v == false {
		t.Error("error")
	}
}
