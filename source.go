package config4live

// Source base implementation get config by key
type Source interface {
	Get(key string) (*Config, error)
}
