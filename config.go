package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hashicorp/hcl"
)

// Config is the structure of the configuration for CLI.
type Config struct {
	StatusOKForNow string
	Verbose        bool
	ConfigFile     string
}

// DefaultConfig returns default structure.
func DefaultConfig() *Config {
	return &Config{
		StatusOKForNow: ":corn: Coverage remained the same at <diff>%",
		Verbose:        false,
		ConfigFile:     ".pipeline",
	}
}

// LoadConfig loads the CLI configuration from conf files.
func LoadConfig(path string) (*Config, error) {
	// Read the HCL file and prepare for parsing
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", path, err)
	}

	// Parse it
	obj, err := hcl.Parse(string(d))
	if err != nil {
		return nil, fmt.Errorf("Error parsing %s: %s", path, err)
	}

	// Build up the result
	var result Config
	if err2 := hcl.DecodeObject(&result, obj); err2 != nil {
		return nil, err2
	}

	return &result, nil
}

// Merge merges other configurations it self.
func (c *config) Merge(otherConfig *Config) *Config {
	if otherConfig.StatusOKForNow != "" {
		c.StatusOKForNow = otherConfig.StatusOKForNow
	}
	c.Verbose = otherConfig.Verbose

	return c
}

// Set sets from Opt
func (c *Config) Set(o Opt) *Config {
	c.Verbose = o.Verbose

	return c
}

// SetFromEnv sets from env variables
func (c *Config) SetFromEnv() *Config {
	token := os.Getenv(strings.ToUpper(Name) + "_ACCESS_TOKEN")
	return c
}
