package removeDuplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	var result int = removeDuplicates([]int{1, 1, 2})
	assert.Equal(t, 2, result)
}

func TestRemoveDuplicates2(t *testing.T) {
	var result int = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	assert.Equal(t, 5, result)
}
