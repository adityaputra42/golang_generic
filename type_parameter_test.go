package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	fmt.Println(total)
	return total
}

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	sumData := Sum[int]([]int{2, 3, 4, 1, 13, 35, 10})
	assert.Equal(t, 68, sumData)
	
	var stringLenght string = Length[string]("main aja")
	assert.Equal(t, "main aja", stringLenght)

	var intLength int = Length[int](200)
	assert.Equal(t, 200, intLength)

}
