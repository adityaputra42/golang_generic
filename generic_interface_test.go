package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(param T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	result := param.GetValue()
	fmt.Println(result)
	return result
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestChangeValue(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Andani")

	assert.Equal(t, "Andani", result)

}
