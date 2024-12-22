package config4live

import "github.com/sayurbox/config4live-go/internal"

// Source base implementation get config by key
type Source interface {
	Get(key string) (*internal.Config, error)
}
