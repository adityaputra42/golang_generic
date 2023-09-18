package golanggeneric

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[X any](bag Bag[X]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestGeneric(t *testing.T) {
	names := Bag[string]{"Paijo", "tukimin"}
	PrintBag(names)
}
