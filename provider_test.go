package config4live

import (
	"errors"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProviderTestSuite struct {
	suite.Suite
	provider *providerImpl
	source   *mockSource
	cache    *cache.Cache
}

// MockStruct is struct type for mocking
type MockStruct struct {
	Name    string
	Id      int
	Address MockAddress
}

// MockAddress is struct type for mocking
type MockAddress struct {
	Postal      int
	Coordinates []float64
}

// MockList is slice type for mocking
type MockList []string

func (s *ProviderTestSuite) SetupTest() {
	s.source = &mockSource{}
	s.provider = NewProvider(
		WithSource(s.source),
	).(*providerImpl)
	s.provider.cache.Flush()
}

func (s *ProviderTestSuite) TestProvider_BindStringFound() {
	s.source.On("Get", "key_string").Return(&Config{
		Name:  "key_string",
		Value: "test_value",
	}, nil)
	actual := s.provider.BindString("key_string", "default_string")
	assert.Equal(s.T(), "test_value", actual)
}

func (s *ProviderTestSuite) TestProvider_BindStringFromCache() {
	s.provider.cache.Set("key_string", "test_cache_value", time.Minute)
	actual := s.provider.BindString("key_string", "default_string")
	assert.Equal(s.T(), "test_cache_value", actual)
}

func (s *ProviderTestSuite) TestProvider_BindStringNotFoundFallbackDefault() {
	s.source.On("Get", "key_string").Return(nil, errors.New("config not found"))
	actual := s.provider.BindString("key_string", "default_string")
	assert.Equal(s.T(), "default_string", actual)
}

func (s *ProviderTestSuite) TestProvider_BindBoolFound() {
	s.source.On("Get", "key_bool").Return(&Config{
		Name:  "key_bool",
		Value: "false",
	}, nil)
	actual := s.provider.BindBool("key_bool", true)
	assert.False(s.T(), actual)
}

func (s *ProviderTestSuite) TestProvider_BindBoolNotFoundFallbackDefault() {
	s.source.On("Get", "key_bool").Return(nil, errors.New("config not found"))
	actual := s.provider.BindBool("key_bool", true)
	assert.True(s.T(), actual)
}

func (s *ProviderTestSuite) TestProvider_BindInt64Found() {
	s.source.On("Get", "key_int").Return(&Config{
		Name:  "key_int",
		Value: "15",
	}, nil)
	actual := s.provider.BindInt64("key_int", 25)
	assert.Equal(s.T(), int64(15), actual)
}

func (s *ProviderTestSuite) TestProvider_BindInt64NotFoundFallbackDefault() {
	s.source.On("Get", "key_int").Return(nil, errors.New("config not found"))
	actual := s.provider.BindInt64("key_int", 25)
	assert.Equal(s.T(), int64(25), actual)
}

func (s *ProviderTestSuite) TestProvider_BindFloat64Found() {
	s.source.On("Get", "key_float").Return(&Config{
		Name:  "key_float",
		Value: "1.5",
	}, nil)
	actual := s.provider.BindFloat64("key_float", 25)
	assert.Equal(s.T(), float64(1.5), actual)
}

func (s *ProviderTestSuite) TestProvider_BindFloat64NotFoundFallbackDefault() {
	s.source.On("Get", "key_float").Return(nil, errors.New("config not found"))
	actual := s.provider.BindFloat64("key_float", 25)
	assert.Equal(s.T(), float64(25), actual)
}

func (s *ProviderTestSuite) TestProvider_BindAnyFound() {
	s.source.On("Get", "key_any_struct").Return(&Config{
		Name:  "key_any_struct",
		Value: "{\"id\":1,\"name\":\"Jhon\",\"address\":{\"postal\":12345,\"coordinates\":[1.32193,0.1299321]}}",
	}, nil)
	s.source.On("Get", "key_any_list").Return(&Config{
		Name:  "key_any_list",
		Value: "[\"ymail.com\",\"test.com\"]",
	}, nil)

	mocked1 := MockStruct{
		Id:   2,
		Name: "Smith",
		Address: MockAddress{
			Postal:      67890,
			Coordinates: []float64{0.88930, 0.32188},
		},
	}
	actual1 := s.provider.BindAny("key_any_struct", mocked1).(MockStruct)
	assert.Equal(s.T(), 1, actual1.Id)
	assert.Equal(s.T(), "Jhon", actual1.Name)
	assert.Equal(s.T(), 12345, actual1.Address.Postal)
	assert.Equal(s.T(), []float64{1.32193, 0.1299321}, actual1.Address.Coordinates)

	mocked2 := MockList{}
	actual2 := s.provider.BindAny("key_any_list", mocked2).(MockList)
	assert.Equal(s.T(), 2, len(actual2))
	assert.Equal(s.T(), MockList{"ymail.com", "test.com"}, actual2)
}

func (s *ProviderTestSuite) TestProvider_BindAnyNotFoundFallbackDefault() {
	s.source.On("Get", "key_any").Return(nil, errors.New("config not found"))

	mocked := MockStruct{
		Id:   2,
		Name: "Smith",
		Address: MockAddress{
			Postal:      67890,
			Coordinates: []float64{0.88930, 0.32188},
		},
	}
	actual := s.provider.BindAny("key_any", mocked).(MockStruct)
	assert.Equal(s.T(), 2, actual.Id)
	assert.Equal(s.T(), "Smith", actual.Name)
	assert.Equal(s.T(), 67890, actual.Address.Postal)
	assert.Equal(s.T(), []float64{0.88930, 0.32188}, actual.Address.Coordinates)
}

func TestProviderSuite(t *testing.T) {
	suite.Run(t, new(ProviderTestSuite))
}

// mock configuration source
type mockSource struct {
	mock.Mock
}

func (m *mockSource) Get(key string) (*Config, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Config), args.Error(1)
}
