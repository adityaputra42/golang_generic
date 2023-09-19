package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func TestStructGeneric(t *testing.T) {

	data := Data[string]{
		First:  "Aditya",
		Second: "Gita",
	}
	fmt.Println(data)

	assert.Equal(t, "Gita", data.ChangeFirst("Gita"))
	assert.Equal(t, "Hello Adit", data.SayHello("Adit"))

}
