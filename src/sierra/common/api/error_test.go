package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {
	parsedErr := NewInternalError(fmt.Errorf("test"), "some internal error")
	assert.Equal(t, 500, parsedErr.httpStatusCode, "code doesn't match")
	assert.Equal(t, "Internal error", parsedErr.UserFacingMessage, "user facing error doesn't match")
}
