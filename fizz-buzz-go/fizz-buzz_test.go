package fizzBuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fizzBuzz_3(t *testing.T) {
	var result []string = fizzBuzz(3)
	var expected []string = []string{"1", "2", "Fizz"}
	assert.Equal(t, result, expected)
}

func Test_fizzBuzz_5(t *testing.T) {
	var result []string = fizzBuzz(5)
	var expected []string = []string{"1", "2", "Fizz", "4", "Buzz"}
	assert.Equal(t, result, expected)
}

func Test_fizzBuzz_15(t *testing.T) {
	var result []string = fizzBuzz(15)
	var expected []string = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	assert.Equal(t, result, expected)
}
