package tools

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Sum(t *testing.T) {
	assert := assert.New(t)

	x := Sum(2, 3)
	assert.Equal(x, 5)
}