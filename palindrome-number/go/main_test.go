package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome_true1(t *testing.T) {
	var result bool = isPalindrome(121)
	assert.Equal(t, result, true)
}

func Test_isPalindrome_true2(t *testing.T) {
	var result bool = isPalindrome(1234321)
	assert.Equal(t, result, true)
}

func Test_isPalindrome_false1(t *testing.T) {
	var result bool = isPalindrome(-121)
	assert.Equal(t, result, false)
}

func Test_isPalindrome_false2(t *testing.T) {
	var result bool = isPalindrome(10)
	assert.Equal(t, result, false)
}
