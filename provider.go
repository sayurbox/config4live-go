package config4live

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/patrickmn/go-cache"
)

type (
	// Provider wrapper to read configuration value
	Provider interface {
		BindString(key string, defaultValue string) string
		BindBool(key string, defaultValue bool) bool
		BindInt64(key string, defaultValue int64) int64
		BindFloat64(key string, defaultValue float64) float64
		BindAny(key string, defaultValue any) any
	}

	providerImpl struct {
		source     Source
		useCache   bool
		expiration time.Duration
		cache      *cache.Cache
	}
)

func (p *providerImpl) BindString(key, defaultValue string) string {
	return p.bind(key, defaultValue)
}

func (p *providerImpl) BindBool(key string, defaultValue bool) bool {
	value, _ := strconv.ParseBool(p.bind(key, defaultValue))
	return value
}

func (p *providerImpl) BindInt64(key string, defaultValue int64) int64 {
	value, _ := strconv.ParseInt(p.bind(key, defaultValue), 10, 64)
	return value
}

func (p *providerImpl) BindFloat64(key string, defaultValue float64) float64 {
	value, _ := strconv.ParseFloat(p.bind(key, defaultValue), 64)
	return value
}

func (p *providerImpl) BindAny(key string, defaultValue any) any {
	encoded := p.bind(key, defaultValue)
	decoded := reflect.New(reflect.TypeOf(defaultValue)).Elem().Interface()

	err := json.Unmarshal([]byte(encoded), &decoded)
	if err != nil {
		return defaultValue
	}

	value := reflect.New(reflect.TypeOf(defaultValue)).Elem().Interface()
	err = mapstructure.Decode(decoded, &value)
	if err != nil {
		return defaultValue
	}

	return value
}

func (p *providerImpl) bind(key string, defaultValue interface{}) string {
	c, found := p.cache.Get(key)
	if found {
		log.Printf("Found %s cache", key)
		return c.(string)
	}
	src, e := p.source.Get(key)
	if e != nil {
		log.Printf("Error get config %s :%s", key, e.Error())
		return fmt.Sprintf("%v", defaultValue)
	}
	p.cache.Set(key, src.Value, p.expiration)
	return src.Value
}

// NewProvider create new configuration provider
func NewProvider(opts ...Option) Provider {
	p := &providerImpl{}
	for _, opt := range opts {
		opt(p)
	}
	p.cache = cache.New(p.expiration, time.Hour)
	return p
}
