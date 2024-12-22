package config4live

import (
	"strings"

	"github.com/sayurbox/config4live-go/grpc"
	"github.com/sayurbox/config4live-go/http"
	"github.com/sayurbox/config4live-go/internal"
)

type SourceProperty struct {
	url          string
	hystrixParam *internal.HystrixParams
}

// SourceOption is mapper of source property
type SourceOption func(s *SourceProperty)

// WithURL set http url
func WithURL(url string) SourceOption {
	return func(s *SourceProperty) {
		s.url = url
	}
}

// WithHystrixTimeout set hystrix timeout
func WithHystrixTimeout(timeout int) SourceOption {
	return func(s *SourceProperty) {
		if s.hystrixParam == nil {
			s.hystrixParam = &internal.HystrixParams{}
		}
		s.hystrixParam.Timeout = timeout
	}
}

// WithHystrixSleepWindow set hystrix sleep_window
func WithHystrixSleepWindow(sleepWindow int) SourceOption {
	return func(s *SourceProperty) {
		if s.hystrixParam == nil {
			s.hystrixParam = &internal.HystrixParams{}
		}
		s.hystrixParam.SleepWindow = sleepWindow
	}
}

// WithHystrixRequestVolumeThreshold set hystrix request_volume_threshold
func WithHystrixRequestVolumeThreshold(requestVolume int) SourceOption {
	return func(s *SourceProperty) {
		if s.hystrixParam == nil {
			s.hystrixParam = &internal.HystrixParams{}
		}
		s.hystrixParam.RequestVolumeThreshold = requestVolume
	}
}

// WithHystrixErrorPercentThreshold set hystrix error_percent_threshold
func WithHystrixErrorPercentThreshold(errorPercent int) SourceOption {
	return func(s *SourceProperty) {
		if s.hystrixParam == nil {
			s.hystrixParam = &internal.HystrixParams{}
		}
		s.hystrixParam.ErrorPercentThreshold = errorPercent
	}
}

// WithHystrixMaxConcurrentRequests set hystrix max_concurrent_requests
func WithHystrixMaxConcurrentRequests(maxRequest int) SourceOption {
	return func(s *SourceProperty) {
		if s.hystrixParam == nil {
			s.hystrixParam = &internal.HystrixParams{}
		}
		s.hystrixParam.MaxConcurrentRequests = maxRequest
	}
}

// WithHystrixCommandName set hystrix command name
func WithHystrixCommandName(commandName string) SourceOption {
	return func(s *SourceProperty) {
		if s.hystrixParam == nil {
			s.hystrixParam = &internal.HystrixParams{}
		}
		s.hystrixParam.Name = commandName
	}
}

// NewSource create configuration source
func NewSource(opts ...SourceOption) Source {
	s := &SourceProperty{}
	for _, opt := range opts {
		opt(s)
	}

	if strings.HasPrefix(s.url, "http") {
		return http.NewHttpSource(http.WithURL(s.url), http.WithHystrix(s.hystrixParam))
	}

	return grpc.NewGrpcSource(grpc.WithURL(s.url), grpc.WithHystrix(s.hystrixParam))
}
