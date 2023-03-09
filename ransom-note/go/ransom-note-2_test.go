package ransomNote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ransomNote2_falseCase1(t *testing.T) {
	var result bool = canConstruct2("a", "b")
	assert.Equal(t, result, false)
}

func Test_ransomNote2_falseCase2(t *testing.T) {
	var result bool = canConstruct2("aa", "ab")
	assert.Equal(t, result, false)
}

func Test_ransomNote2_falseCase1_trueCase(t *testing.T) {
	var result bool = canConstruct2("aa", "aab")
	assert.Equal(t, result, true)
}

func Test_ransomNote2_trueCase2(t *testing.T) {
	var result bool = canConstruct2("aab", "baa")
    assert.Equal(t, result, true)
}
