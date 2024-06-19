package configs

import "os"

// DB ...
type DB struct {
	Username string
	Password string
	Port     int
	Database string
	Host     string
}

// Flags contains the command line flags.
type Flags struct {
	Addr string
	DB   *DB
}

// New ...
func New() *Config {
	return &Config{
		Flags: &Flags{
			DB: &DB{
				Username: "example",
				Password: "example",
				Port:     5432,
				Database: "example",
				Host:     "host.docker.internal",
			},
		},
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
