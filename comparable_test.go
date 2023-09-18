package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CompareData[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		fmt.Println(value1, "==", value2)
		return true
	} else {
		fmt.Println(value1, "!=", value2)
		return false
	}
}

func SumNumbers1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestCompareData(t *testing.T) {
	CompareData[string]("Saya", "Saya")
	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers2(ints),
		SumNumbers2(floats))
	assert.True(t, CompareData[string]("Saya", "Saya"))
	assert.True(t, CompareData[int](100, 100))
	assert.True(t, CompareData[bool](true, true))

}
