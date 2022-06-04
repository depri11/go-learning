package generics

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
	var result string = Length[string]("dev")
	assert.Equal(t, "dev", result)

	var resultNumber = Length[int](10)
	assert.Equal(t, 10, resultNumber)
}

func MultipleParameter[T1 any, T2 any](param T1, param2 T2) {
	fmt.Println(param)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("dev", 10)
	MultipleParameter[int, string](10, "dev")
}
