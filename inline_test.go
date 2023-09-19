package golanggeneric

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int32 | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Aditya", "Putra", "Pratama",
	}
	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Aditya", first)

}

func TestInline(t *testing.T) {
	assert.Equal(t, int32(100), FindMin[int32](100, 100))
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, 100.0, FindMin(100.0, 100.0))

}
