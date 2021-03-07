package main

import (
	"fmt"

	"hello_19/common"
)

func main() {

	key := "cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4="
	acc := "62174378490347"

	s := common.Encrypt(key, acc)

	fmt.Println(s)

	x := common.Decrypt(key, s)

	fmt.Println(x)
}
