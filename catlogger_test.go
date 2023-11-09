package catlogger

import "testing"
import "github.com/stretchr/testify/assert"

func TestLogger(t *testing.T) {
	logger := MakeLogger("foo,bar", "hello", false)
	assert.NotNil(t, logger, "could not make logger")
	assert.True(t, logger.HasCategory("foo"), "logger lacks category foo")
	assert.True(t, logger.HasCategory("bar"), "logger lacks category bar")
	assert.True(t, !logger.HasCategory("baz"), "logger has category baz")
}
