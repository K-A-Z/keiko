package keiko

import (
	"testing"

//	"fmt"

)

type lampBitTest struct {
	color string
	state int
	out   string
	err   string
}

var lampBitTests = []lampBitTest{
	{"red", 0, "0xxxxxxx", ""},
	{"red", 1, "1xxxxxxx", ""},
	{"red", 2, "2xxxxxxx", ""},
	{"red", 3, "3xxxxxxx", ""},
	{"yellow", 0, "x0xxxxxx", ""},
	{"yellow", 1, "x1xxxxxx", ""},
	{"yellow", 2, "x2xxxxxx", ""},
	{"yellow", 3, "x3xxxxxx", ""},
	{"green", 0, "xx0xxxxx", ""},
	{"green", 1, "xx1xxxxx", ""},
	{"green", 2, "xx2xxxxx", ""},
	{"green", 3, "xx3xxxxx", ""},
	//	{"black",0,"","Unexpected lamp color or state"},
}

func TestSetLampStatusBit(t *testing.T) {

	for _, lbt := range lampBitTests {
		result, err := setLampStatusBit(lbt.color, lbt.state)
		if result != lbt.out || err != lbt.err {
			t.Errorf("lamp bit wrong!!" + err.Error())
		}
	}

}
