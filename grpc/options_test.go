package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrpcOptions(t *testing.T) {
	s := NewGrpcSource(WithURL("url"))
	assert.Equal(t, "url", s.url)
	s = NewGrpcSource(WithHystrixCommandName(""))
	assert.NotNil(t, s.hystrixParam)
	assert.Equal(t, "live-config-command-key", s.hystrixParam.Name)
	s = NewGrpcSource(WithHystrixTimeout(10))
	assert.Equal(t, 10, s.hystrixParam.Timeout)
	s = NewGrpcSource(WithHystrixErrorPercentThreshold(5))
	assert.Equal(t, 5, s.hystrixParam.ErrorPercentThreshold)
	s = NewGrpcSource(WithHystrixMaxConcurrentRequests(15))
	assert.Equal(t, 15, s.hystrixParam.MaxConcurrentRequests)
	s = NewGrpcSource(WithHystrixRequestVolumeThreshold(20))
	assert.Equal(t, 20, s.hystrixParam.RequestVolumeThreshold)
	s = NewGrpcSource(WithHystrixSleepWindow(100))
	assert.Equal(t, 100, s.hystrixParam.SleepWindow)
}
