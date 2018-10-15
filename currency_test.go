package eth_utils

import (
	"testing"
)

func Test_From_Wei_1(t *testing.T) {
	value, err := FromWei(1000000000000000000, "ether")
	if err != nil{
		t.Error("error")
	}
	if value != 1{
		t.Error("error")
	}
}

func Test_From_Wei_2(t *testing.T) {
	value, err := FromWei(1000000000000000000, "eether")
	if err == nil {
		t.Error("error")
	}
	if value != 0{
		t.Error("error")
	}
}

