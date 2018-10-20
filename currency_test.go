package eth_utils

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestFromWei1(t *testing.T) {
	arg, _ := decimal.NewFromString("1000000000000000000")
	value, err := FromWei(arg, "ether")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1")
	if value.Equal(v) == false {
		t.Error("error")
	}
}

func TestFromWei2(t *testing.T) {
	arg, _ := decimal.NewFromString("1000000000000000000")
	value, err := FromWei(arg, "tether")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("0.000000000001")
	if value.Equal(v) == false {
		t.Error("error")
	}
}

func Test_From_Wei_3(t *testing.T) {
	arg, _ := decimal.NewFromString("100")
	value, err := FromWei(arg, "hoge")
	if err == nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("0")
	if value.Equal(v) == false {
		t.Error("error")
	}
}

func TestGetEtherUnit1(t *testing.T) {
	unit, err := getEtherUnit("ether")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1000000000000000000")
	if unit.Equal(v) == false {
		t.Error("error")
	}
}

func TestGetEtherUnit2(t *testing.T) {
	unit, err := getEtherUnit("tether")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1000000000000000000000000000000")
	if unit.Equal(v) == false {
		t.Error("error")
	}
}

func TestGetEtherUnit3(t *testing.T) {
	unit, err := getEtherUnit("wei")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1")
	if unit.Equal(v) == false {
		t.Error("error")
	}
}

func TestGetEtherUnitError(t *testing.T) {
	unit, err := getEtherUnit("hoge")
	if err == nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("0")
	if unit.Equal(v) == false {
		t.Error("error")
	}
}

func TestToWei1(t *testing.T) {
	arg, _ := decimal.NewFromString("1")
	value, err := ToWei(arg, "ether")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1000000000000000000")
	if value.Equal(v) == false {
		t.Error("error")
	}
}

func TestToWei2(t *testing.T) {
	arg, _ := decimal.NewFromString("1")
	value, err := ToWei(arg, "wei")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1")
	if value.Equal(v) == false {
		t.Error("error")
	}
}

func TestToWei3(t *testing.T) {
	arg, _ := decimal.NewFromString("1")
	value, err := ToWei(arg, "tether")
	if err != nil {
		t.Error("error")
	}
	v, _ := decimal.NewFromString("1000000000000000000000000000000")
	if value.Equal(v) == false {
		t.Error("error")
	}
}
