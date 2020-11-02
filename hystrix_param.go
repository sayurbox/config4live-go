package config4live

import "github.com/afex/hystrix-go/hystrix"

// HystrixParams global configuration for hystrix
type HystrixParams struct {
	Name                   string
	Timeout                int
	SleepWindow            int
	RequestVolumeThreshold int
	ErrorPercentThreshold  int
	MaxConcurrentRequests  int
}

// HystrixConfig create command config
func (h *HystrixParams) HystrixConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                h.Timeout,
		SleepWindow:            h.SleepWindow,
		RequestVolumeThreshold: h.RequestVolumeThreshold,
		ErrorPercentThreshold:  h.ErrorPercentThreshold,
		MaxConcurrentRequests:  h.MaxConcurrentRequests,
	}
}
