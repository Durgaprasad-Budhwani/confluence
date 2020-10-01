package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleted(t *testing.T) {
	assert := assert.New(t)
	buf, err := Parse(`-Hello-`)
	assert.NoError(err)
	assert.Equal("<s>Hello</s>", buf)
	buf, err = Parse("-Hello-\n")
	assert.NoError(err)
	assert.Equal("<s>Hello</s>\n", buf)
}
