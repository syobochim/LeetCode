package maxProfit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit1(t *testing.T) {
	var result int = maxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 7, result)
}

func TestMaxProfit2(t *testing.T) {
	var result int = maxProfit([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 4, result)
}

func TestMaxProfit3(t *testing.T) {
	var result int = maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, 0, result)
}
