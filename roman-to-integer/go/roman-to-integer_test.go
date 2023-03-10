package roman;

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInteger_III(t *testing.T) {
    result := romanToInt("III")
	assert.Equal(t, 3, result)
}

func Test_romanToInteger_LVIII(t *testing.T) {
    result := romanToInt("LVIII")
    assert.Equal(t, 58, result)
}

func Test_romanToInteger_MCMXCIV(t *testing.T) {
    result := romanToInt("MCMXCIV")
    assert.Equal(t, 1994, result)
}