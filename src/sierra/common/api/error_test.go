package api_test

import (
	"fmt"
	"sierra/common/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {
	aErr := api.NewInternalError(fmt.Errorf("test"), "")

	parsedErr := api.NewInternalError(aErr, "")
	assert.Equal(t, aErr.Code, parsedErr.Code, "code doesn't match")
	assert.Equal(t, aErr.Err, parsedErr.Err, "err doesn't match")
}

func TestNonAPIError(t *testing.T) {
	err := fmt.Errorf("test")

	parsedErr := api.NewInternalError(err, "")
	assert.Equal(t, 0, parsedErr.Code, "code doesn't match")
}
