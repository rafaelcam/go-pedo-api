package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenItIsMultipleOfThree(t *testing.T) {
	expected := []string{"Pé"}
	result, _ := GetPeDoByRange(3, 3)

	assert.Equal(t,1, len(result))
	assert.Equal(t, expected, result)
}

func TestWhenItIsMultipleOfFive(t *testing.T) {
	expected := []string{"Do"}
	result, _ := GetPeDoByRange(5, 5)

	assert.Equal(t, expected, result)
}

func TestWhenItIsMultipleOfThreeAndFive(t *testing.T) {
	expected := []string{"PéDo"}
	result, _ := GetPeDoByRange(75, 75)

	assert.Equal(t, expected, result)
}

func TestWhenItIsNotMultipleOfThreeOrFive(t *testing.T) {
	expected := []string{"1"}
	result, _ := GetPeDoByRange(1, 1)

	assert.Equal(t, expected, result)
}

func TestWhenStartIsOneAndEndIsFifteen(t *testing.T) {
	expected := []string{"1","2","Pé","4","Do","Pé","7","8","Pé","Do","11","Pé","13","14","PéDo"}
	result, _ := GetPeDoByRange(1, 15)

	assert.Equal(t, expected, result)
}

func TestWhenStartOrEndIsNegative(t *testing.T) {
	result, err := GetPeDoByRange(-1, 5)

	assert.Nil(t, result)
	assert.Equal(t, err.Error(), "Only positive values ​​are accepted.")
}

func TestWhenStartGreaterThanTheEndValue(t *testing.T) {
	result, err := GetPeDoByRange(5, 3)

	assert.Nil(t, result)
	assert.Equal(t, err.Error(), "Start value can not be greater than the End value.")
}