package config4live

import "time"

// Option the provider options
type Option func(p *providerImpl)

// WithSource set provider source
func WithSource(source Source) Option {
	return func(p *providerImpl) {
		p.source = source
	}
}

// WithCache use inmemory cache
func WithCache(useCache bool) Option {
	return func(p *providerImpl) {
		p.useCache = useCache
	}
}

// WithExpiration cache expiration
func WithExpiration(ttl time.Duration) Option {
	return func(p *providerImpl) {
		p.expiration = ttl
	}
}
