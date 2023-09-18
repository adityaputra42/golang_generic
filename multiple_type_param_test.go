package golanggeneric

import (
	"fmt"
	"testing"
)

func MultipleParam[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}
func TestMultipleParam(t *testing.T) {

	MultipleParam[string, int]("Paijo", 100)
	MultipleParam[bool, float32](true, 100)

}
