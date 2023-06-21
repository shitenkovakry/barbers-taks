package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	countOfMenInCity := 9000
	countOfBarbersInBarbershop := 57

	expectedEnoughBarbers, expectedReason := true, "барберы не нужны"
	actualEnoughBarbers, actualReason := CalculateIfThereAreEnoughBarbers(countOfBarbersInBarbershop, countOfMenInCity)

	assert.Equal(t, expectedEnoughBarbers, actualEnoughBarbers)
	assert.Equal(t, expectedReason, actualReason)
}
