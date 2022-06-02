package problemsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLongSubarrays(t *testing.T) {
	assert.Equal(t, 0, countLongSubarrays([]int{}))
	assert.Equal(t, 2, countLongSubarrays([]int{1, 3, 4, 2, 7, 5, 6, 9, 8}))
	assert.Equal(t, 9, countLongSubarrays([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
