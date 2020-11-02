package grpc

import (
	"github.com/sayurbox/config4live-go"
)

// Option the grpc source options
type Option func(s *GrpcSource)

// WithURL set grpc url
func WithURL(url string) Option {
	return func(s *GrpcSource) {
		s.url = url
	}
}

// WithHystrixTimeout set hystrix timeout
func WithHystrixTimeout(timeout int) Option {
	return func(s *GrpcSource) {
		if s.hystrixParam == nil {
			s.hystrixParam = &config4live.HystrixParams{}
		}
		s.hystrixParam.Timeout = timeout
	}
}

// WithHystrixSleepWindow set hystrix sleep_window
func WithHystrixSleepWindow(sleepWindow int) Option {
	return func(s *GrpcSource) {
		if s.hystrixParam == nil {
			s.hystrixParam = &config4live.HystrixParams{}
		}
		s.hystrixParam.SleepWindow = sleepWindow
	}
}

// WithHystrixRequestVolumeThreshold set hystrix request_volume_threshold
func WithHystrixRequestVolumeThreshold(requestVolume int) Option {
	return func(s *GrpcSource) {
		if s.hystrixParam == nil {
			s.hystrixParam = &config4live.HystrixParams{}
		}
		s.hystrixParam.RequestVolumeThreshold = requestVolume
	}
}

// WithHystrixErrorPercentThreshold set hystrix error_percent_threshold
func WithHystrixErrorPercentThreshold(errorPercent int) Option {
	return func(s *GrpcSource) {
		if s.hystrixParam == nil {
			s.hystrixParam = &config4live.HystrixParams{}
		}
		s.hystrixParam.ErrorPercentThreshold = errorPercent
	}
}

// WithHystrixMaxConcurrentRequests set hystrix max_concurrent_requests
func WithHystrixMaxConcurrentRequests(maxRequest int) Option {
	return func(s *GrpcSource) {
		if s.hystrixParam == nil {
			s.hystrixParam = &config4live.HystrixParams{}
		}
		s.hystrixParam.MaxConcurrentRequests = maxRequest
	}
}

// WithHystrixCommandName set hystrix command name
func WithHystrixCommandName(commandName string) Option {
	return func(s *GrpcSource) {
		if s.hystrixParam == nil {
			s.hystrixParam = &config4live.HystrixParams{}
		}
		s.hystrixParam.Name = commandName
	}
}
