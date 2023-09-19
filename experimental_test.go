package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 100.0))

}

func TestExperimentalMap(t *testing.T) {
	first := map[string]string{
		"Name": "Aditya",
	}
	second := map[string]string{
		"Name": "Aditya",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"Aditya"}
	second := []string{"Aditya"}

	assert.True(t, slices.Equal(first, second))
}
