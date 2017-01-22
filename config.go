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
	APIEndpoint    string
	AccessToken    string
	Comment        bool
	Verbose        bool
	ConfigFile     string
}

// DefaultConfig returns default structure.
func DefaultConfig() *Config {
	return &Config{
		StatusOKForNow: ":corn: Coverage remained the same at <diff>%",
		APIEndpoint:    "https://api.github.com/",
		AccessToken:    "",
		Comment:        false,
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
	if otherConfig.APIEndpoint != "" {
		c.APIEndpoint = otherConfig.APIEndpoint
	}
	if otherConfig.AccessToken != "" {
		c.AccessToken = otherConfig.AccessToken
	}
	c.Verbose = otherConfig.Verbose
	c.Comment = otherConfig.Comment

	return c
}

// Set sets from Opt
func (c *Config) Set(o Opt) *Config {
	if o.APIEndpoint != "" {
		c.APIEndpoint = o.APIEndpoint
	}
	if o.AccessToken != "" {
		c.AccessToken = o.AccessToken
	}
	c.Verbose = o.Verbose
	c.Comment = o.Comment

	return c
}

// SetFromEnv sets from env variables
func (c *Config) SetFromEnv() *Config {
	token := os.Getenv(strings.ToUpper(Name) + "_ACCESS_TOKEN")
	if token != "" {
		c.AccessToken = token
	}
	return c
}
