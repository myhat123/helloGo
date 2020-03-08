package tools

import "testing"

func Test_Sum(t *testing.T) {
	x := Sum(2, 3)
	expect := 6
	if x != expect {
		t.Errorf("got [%d] expected [%d]", x, expect)
	} 
}