package embeddedpostgres

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

// Config maintains the runtime configuration for the Postgres process to be created.
type Config struct {
	version             PostgresVersion
	port                uint32
	database            string
	username            string
	password            string
	runtimePath         string
	dataPath            string
	locale              string
	startTimeout        time.Duration
	logger              io.Writer
	cacheLocator        CacheLocator
	remoteFetchStrategy RemoteFetchStrategy
	unpacker            Unpacker
}

// DefaultConfig provides a default set of configuration to be used "as is" or modified using the provided builders.
// The following can be assumed as defaults:
// Version:      12
// Port:         5432
// Database:     postgres
// Username:     postgres
// Password:     postgres
// StartTimeout: 15 Seconds
func DefaultConfig() Config {
	return Config{
		version:      V12,
		port:         5432,
		database:     "postgres",
		username:     "postgres",
		password:     "postgres",
		startTimeout: 15 * time.Second,
		logger:       os.Stdout,
	}
}

// Version will set the Postgres binary version.
func (c Config) Version(version PostgresVersion) Config {
	c.version = version
	return c
}

// Port sets the runtime port that Postgres can be accessed on.
func (c Config) Port(port uint32) Config {
	c.port = port
	return c
}

// Database sets the database name that will be created.
func (c Config) Database(database string) Config {
	c.database = database
	return c
}

// Username sets the username that will be used to connect.
func (c Config) Username(username string) Config {
	c.username = username
	return c
}

// Password sets the password that will be used to connect.
func (c Config) Password(password string) Config {
	c.password = password
	return c
}

// RuntimePath sets the path that will be used for the extracted Postgres runtime directory.
// If Postgres data directory is not set with DataPath(), this directory is also used as data directory.
func (c Config) RuntimePath(path string) Config {
	if filepath.IsAbs(path) {
		c.runtimePath = path
	} else {
		wd, _ := os.Getwd()
		c.runtimePath = filepath.Join(wd, path)
	}
	return c
}

// DataPath sets the path that will be used for the Postgres data directory.
// If this option is set, a previously initialized data directory will be reused if possible.
func (c Config) DataPath(path string) Config {
	if filepath.IsAbs(path) {
		c.dataPath = path
	} else {
		wd, _ := os.Getwd()
		c.dataPath = filepath.Join(wd, path)
	}
	return c
}

// Locale sets the default locale for initdb
func (c Config) Locale(locale string) Config {
	c.locale = locale
	return c
}

// StartTimeout sets the max timeout that will be used when starting the Postgres process and creating the initial database.
func (c Config) StartTimeout(timeout time.Duration) Config {
	c.startTimeout = timeout
	return c
}

// Logger sets the logger for postgres output
func (c Config) Logger(logger io.Writer) Config {
	c.logger = logger
	return c
}

// CacheLocator sets the cache locator function
func (c Config) CacheLocator(cl CacheLocator) Config {
	c.cacheLocator = cl
	return c
}

// Unpacker sets the unpacker function
func (c Config) Unpacker(u Unpacker) Config {
	c.unpacker = u
	return c
}

// PostgresVersion represents the semantic version used to fetch and run the Postgres process.
type PostgresVersion string

// Predefined supported Postgres versions.
const (
	V13 = PostgresVersion("13.2.0")
	V12 = PostgresVersion("12.6.0")
	V11 = PostgresVersion("11.11.0")
	V10 = PostgresVersion("10.16.0")
	V9  = PostgresVersion("9.6.21")
)
