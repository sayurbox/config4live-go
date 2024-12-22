package internal

// Config configuration property
type Config struct {
	ID          string
	Name        string
	Value       string
	Description string
	Owner       string
}

// GetConfigResponse response wrapper Config and error
type GetConfigResponse struct {
	Output *Config
	Error  error
}
