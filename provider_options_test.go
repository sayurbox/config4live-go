package config4live

import (
	"testing"
	"time"

	"github.com/sayurbox/config4live-go/internal"
	"github.com/stretchr/testify/assert"
)

type objectSource struct {
}

func (o *objectSource) Get(key string) (*internal.Config, error) {
	return nil, nil
}

func TestProviderOptions(t *testing.T) {
	p := NewProvider(
		WithSource(&objectSource{}),
		WithCache(true),
		WithExpiration(5*time.Second),
	)
	assert.NotNil(t, p)
	assert.NotNil(t, p.(*providerImpl).source)
	assert.True(t, p.(*providerImpl).useCache)
	assert.Equal(t, 5*time.Second, p.(*providerImpl).expiration)
}
