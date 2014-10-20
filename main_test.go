package fixedlen

import "fmt"
import "testing"

type TestType struct {
	X    int    `fw:"8"`
	Y    int    `fw:"8"`
	name string `fw:"20"`
}

func TestMain(t *testing.T) {
	r := TestType{1, 2, "test"}
	fmt.Println(Marshal(r))
}
