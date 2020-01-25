package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}

func TestStatusNotDown(t *testing.T) {
	status := "waiting"
	assert.NotEqual(t, status, "down")
}

func TestCalculate_TableDrive(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}