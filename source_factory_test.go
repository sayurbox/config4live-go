package config4live

import (
	"testing"

	"github.com/sayurbox/config4live-go/grpc"
	"github.com/sayurbox/config4live-go/http"
	"github.com/stretchr/testify/assert"
)

func TestWithURL(t *testing.T) {
	sp := &SourceProperty{}
	WithURL("http://localhost")(sp)
	assert.Equal(t, "http://localhost", sp.url)
}

func TestWithHystrixTimeout(t *testing.T) {
	sp := &SourceProperty{}
	WithHystrixTimeout(5000)(sp)
	assert.Equal(t, 5000, sp.hystrixParam.Timeout)
}

func TestWithHystrixSleepWindow(t *testing.T) {
	sp := &SourceProperty{}
	WithHystrixSleepWindow(1000)(sp)
	assert.Equal(t, 1000, sp.hystrixParam.SleepWindow)
}

func TestWithHystrixRequestVolumeThreshold(t *testing.T) {
	sp := &SourceProperty{}
	WithHystrixRequestVolumeThreshold(20)(sp)
	assert.Equal(t, 20, sp.hystrixParam.RequestVolumeThreshold)
}

func TestWithHystrixErrorPercentThreshold(t *testing.T) {
	sp := &SourceProperty{}
	WithHystrixErrorPercentThreshold(30)(sp)
	assert.Equal(t, 30, sp.hystrixParam.ErrorPercentThreshold)
}

func TestWithHystrixMaxConcurrentRequests(t *testing.T) {
	sp := &SourceProperty{}
	WithHystrixMaxConcurrentRequests(10)(sp)
	assert.Equal(t, 10, sp.hystrixParam.MaxConcurrentRequests)
}

func TestWithHystrixCommandName(t *testing.T) {
	sp := &SourceProperty{}
	WithHystrixCommandName("test-command")(sp)
	assert.Equal(t, "test-command", sp.hystrixParam.Name)
}

func TestNewSource_Http(t *testing.T) {
	source := NewSource(WithURL("http://localhost"))
	_, ok := source.(*http.HttpSource)
	assert.True(t, ok)
}

func TestNewSource_Grpc(t *testing.T) {
	source := NewSource(WithURL("grpc://localhost"))
	_, ok := source.(*grpc.GrpcSource)
	assert.True(t, ok)
}
