package fixedlen

import "testing"

type TestType struct {
	X    int     `fw:"3"`
	Y    uint    `fw:"8"`
	Z    float32 `fw:"20"`
	name string  `fw:"20"`
}

func TestMain(t *testing.T) {
	expected := "-1 2       -45.3234            test                "
	testStruct := TestType{-1, 2, -45.3234, "test"}
	result := Encode(testStruct)
	if string(result) != expected {
		t.Errorf("Not encoded as expected.\nGot: %s\nExpected: %s\n",
			result, expected)
	}
}

func TestPadString(t *testing.T) {
	input := "a"
	output := padString(input, 10)
	if output != "a         " {
		t.Errorf("String was not properly padded: '%s'\n", output)
	}
}
