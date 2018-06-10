package validators

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWhenStartIsNegative(t *testing.T) {
	err := ValidateRange(-1, 100)
	assert.Equal(t, err.Error(), "Only positive values ​​are accepted.")
}

func TestWhenEndIsNegative(t *testing.T) {
	err := ValidateRange(2, -6)
	assert.Equal(t, err.Error(), "Only positive values ​​are accepted.")
}

func TestWhenIsValid(t *testing.T) {
	err := ValidateRange(2, 5)
	assert.Nil(t, err)
}