package eth_utils

import (
	"testing"
	"eth-utils"
)

func Test_From_Wei_1(t *testing.T) {
	w, e := FromWei(1, "ether")
	if e != nil{
		t.Error("除算関数のテストが通りません") // もし予定されたものでなければエラーを発生させます。
	}
	fmt.Printf("", w)
}
