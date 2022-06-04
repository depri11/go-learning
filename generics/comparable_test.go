package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("devri", "devri"))
	assert.True(t, IsSame[int](10, 10))
	assert.True(t, IsSame[bool](true, true))
}
