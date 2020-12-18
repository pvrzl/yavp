package yavp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAfterBefore(t *testing.T) {
	assert.Equal(t, ErrInvalidValue, IsAfter(time.Now()).Validate(time.Now().Add(-1*time.Hour)), invalidTestComment)
	assert.NoError(t, IsAfter(time.Now()).Validate(time.Now().Add(1*time.Hour)), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsBefore(time.Now()).Validate(time.Now().Add(1*time.Hour)), invalidTestComment)
	assert.NoError(t, IsBefore(time.Now()).Validate(time.Now().Add(-1*time.Hour)), validTestComment)
}
