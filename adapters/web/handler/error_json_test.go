package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello json error"

	result := jsonError(msg)

	require.Equal(t, `{"message":"Hello json error"}`, string(result))
}
