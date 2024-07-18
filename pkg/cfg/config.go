package cfg

import (
	"os"
)

// Flags contains the command line flags.
type Flags struct {
	Addr                    string `envconfig:"KNOX_ADDR" default:":8084"`
	DatabaseURI             string `envconfig:"KNOX_DATABASE_URI" default:"postgres://root@host.docker.internal:26257/defaultdb?sslmode=disable"`
	DatabaseTablePrefix     string `envconfig:"KNOX_DATABASE_TABLE_PREFIX" default:"knox_"`
	FGAApiUrl               string `envconfig:"KNOX_FGA_API_URL" default:"http://host.docker.internal:8080"`
	FGAStoreID              string `envconfig:"KNOX_FGA_STORE_ID" default:""`
	FGAAuthorizationModelID string `envconfig:"KNOX_FGA_AUTHORIZATION_MODEL_ID" default:""`
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{}
}

// New ...
func New() *Config {
	return &Config{
		Flags: NewFlags(),
	}
}

// Config ...
type Config struct {
	Flags *Flags
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
