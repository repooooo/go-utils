// Package loader github.com/repooooo/go-utils/loader/loader.go
package loader

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// Loader - type responsible for loading configuration.
type Loader struct {
	configPath string
}

// NewConfigLoader - constructor for the config loader.
func NewConfigLoader(configPath string) *Loader {
	return &Loader{
		configPath: configPath,
	}
}

// MustLoad - loads the configuration from the file, with the path provided in the arguments.
// The user must provide their custom configuration structure.
//
// Example usage:
//
//	 cfg := &MyConfig{}
//		configPath := loader.FetchConfigPath()
//		cl := loader.NewConfigLoader(configPath)
//		cl.MustLoad(cfg)
func (cl *Loader) MustLoad(cfg interface{}) interface{} {
	if cl.configPath == "" {
		panic("config path is empty")
	}

	return cl.MustLoadPath(cfg)
}

// MustLoadPath - loads the configuration from the given path.
// This function uses a pointer to an arbitrary structure.
func (cl *Loader) MustLoadPath(cfg interface{}) interface{} {
	// check if file exists
	if _, err := os.Stat(cl.configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + cl.configPath)
	}

	// Read the configuration into the provided structure
	if err := cleanenv.ReadConfig(cl.configPath, cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return cfg
}

// FetchConfigPath - fetches the config path from command-line flag or environment variable.
// Priority: flag > env > default.
func FetchConfigPath() string {
	var res string

	// Parse command-line flag for config file path
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	// If no flag is provided, try fetching from environment variable
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
