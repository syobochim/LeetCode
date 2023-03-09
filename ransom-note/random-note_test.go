package ransomNote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ransomNote_falseCase1(t *testing.T) {
	var result bool = canConstruct("a", "b")
	assert.Equal(t, result, false)
}

func Test_ransomNote_falseCase2(t *testing.T) {
	var result bool = canConstruct("aa", "ab")
	assert.Equal(t, result, false)
}

func Test_ransomNote_trueCase(t *testing.T) {
	var result bool = canConstruct("aa", "aab")
	assert.Equal(t, result, true)
}

func Test_ransomNote_trueCase2(t *testing.T) {
	var result bool = canConstruct("aab", "baa")
	assert.Equal(t, result, true)
}
