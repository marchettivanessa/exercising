package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpingOnCloudSuccess(t *testing.T) {
	testName := "Function is called and returns the correct minimum number of times the loop iterates over value 0, jumping 1 os 2 steps tops"
	t.Run(testName, func(t *testing.T) {
		testArrayCloud := []int{0, 0, 1, 0, 1, 0}
		expectedResult := 3
		err := jumpingOnClouds(testArrayCloud)

		assert.Equal(t, expectedResult, err)

	})

}

func TestJumpingOnCloudError(t *testing.T) {
	testName := "Function is called and returns the wrong minimum number of times the loop iterates over value 0, jumping 1 os 2 steps tops"
	t.Run(testName, func(t *testing.T) {
		testArrayCloud := []int{0, 0, 1, 0, 1, 0}
		expectedResult := 5
		err := jumpingOnClouds(testArrayCloud)

		assert.NotEqual(t, expectedResult, err)

	})

}
