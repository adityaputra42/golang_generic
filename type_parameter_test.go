package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)



func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {

	
	var stringLenght string = Length[string]("main aja")
	assert.Equal(t, "main aja", stringLenght)

	var intLength int = Length[int](200)
	assert.Equal(t, 200, intLength)

}
