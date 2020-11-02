package config4live

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHystixParams(t *testing.T) {
	h := &HystrixParams{
		Timeout:                2000,
		ErrorPercentThreshold:  25,
		MaxConcurrentRequests:  5,
		RequestVolumeThreshold: 10,
		SleepWindow:            400,
	}
	assert.NotNil(t, h)
	cmd := h.HystrixConfig()
	assert.NotNil(t, cmd)
	assert.Equal(t, 2000, cmd.Timeout)
	assert.Equal(t, 25, cmd.ErrorPercentThreshold)
	assert.Equal(t, 5, cmd.MaxConcurrentRequests)
	assert.Equal(t, 10, cmd.RequestVolumeThreshold)
	assert.Equal(t, 400, cmd.SleepWindow)
}
