package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrpcOptions(t *testing.T) {
	s := NewHttpSource(WithURL("url"))
	assert.Equal(t, "url", s.url)
	s = NewHttpSource(WithHystrixCommandName(""))
	assert.NotNil(t, s.hystrixParam)
	assert.Equal(t, "live-config-command-key", s.hystrixParam.Name)
	s = NewHttpSource(WithHystrixTimeout(10))
	assert.Equal(t, 10, s.hystrixParam.Timeout)
	s = NewHttpSource(WithHystrixErrorPercentThreshold(5))
	assert.Equal(t, 5, s.hystrixParam.ErrorPercentThreshold)
	s = NewHttpSource(WithHystrixMaxConcurrentRequests(15))
	assert.Equal(t, 15, s.hystrixParam.MaxConcurrentRequests)
	s = NewHttpSource(WithHystrixRequestVolumeThreshold(20))
	assert.Equal(t, 20, s.hystrixParam.RequestVolumeThreshold)
	s = NewHttpSource(WithHystrixSleepWindow(100))
	assert.Equal(t, 100, s.hystrixParam.SleepWindow)
}
