package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EvalPostfix(t *testing.T) {
	assert.Equal(t, 10, EvalPostfix("232+*"))
	assert.Equal(t, 16, EvalPostfix("842/*"))
	assert.Equal(t, 11, EvalPostfix("245++"))
	assert.Equal(t, 2, EvalPostfix("552--"))
	assert.Equal(t, -5, EvalPostfix("252+-"))
	assert.Equal(t, 5, EvalPostfix("252-+"))
	assert.Equal(t, 11, EvalPostfix("9696-/+"))
}
