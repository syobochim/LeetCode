package roman;

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_romanToInteger2_III(t *testing.T) {
    result := romanToInt2("III")
    assert.Equal(t, 3, result)
}

func Test_romanToInteger2_LVIII(t *testing.T) {
    result := romanToInt2("LVIII")
    assert.Equal(t, 58, result)
}

func Test_romanToInteger2_MCMXCIV(t *testing.T) {
    result := romanToInt2("MCMXCIV")
    assert.Equal(t, 1994, result)
}