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
	TestCMD        string
	CoverageCMD    string
	BaseBranch     string
	StatusName     string
	StatusOK       string
	StatusNG       string
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
		TestCMD:        "",
		CoverageCMD:    "",
		BaseBranch:     "master",
		StatusName:     "coverage/cos",
		StatusOK:       ":cake: Coverage increased (+<diff>%) to <coverage>%",
		StatusNG:       ":jack_o_lantern: Coverage decreased (-<diff>%) to <coverage>%",
		StatusOKForNow: ":corn: Coverage remained the same at <diff>%",
		APIEndpoint:    "https://api.github.com/",
		AccessToken:    "",
		Comment:        false,
		Verbose:        false,
		ConfigFile:     ".cos",
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
	if otherConfig.TestCMD != "" {
		c.TestCMD = otherConfig.TestCMD
	}
	if otherConfig.CoverageCMD != "" {
		c.CoverageCMD = otherConfig.CoverageCMD
	}
	if otherConfig.BaseBranch != "" {
		c.BaseBranch = otherConfig.BaseBranch
	}
	if otherConfig.StatusName != "" {
		c.StatusName = otherConfig.StatusName
	}
	if otherConfig.StatusOK != "" {
		c.StatusOK = otherConfig.StatusOK
	}
	if otherConfig.StatusOKForNow != "" {
		c.StatusOKForNow = otherConfig.StatusOKForNow
	}
	if otherConfig.StatusNG != "" {
		c.StatusNG = otherConfig.StatusNG
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

// Set sets from Ops
func (c *Config) Set(o Ops) *Config {
	if o.TestCMD != "" {
		c.TestCMD = o.TestCMD
	}
	if o.CoverageCMD != "" {
		c.CoverageCMD = o.CoverageCMD
	}
	if o.BaseBranch != "" {
		c.BaseBranch = o.BaseBranch
	}
	if o.StatusName != "" {
		c.StatusName = o.StatusName
	}
	if o.StatusOK != "" {
		c.StatusOK = o.StatusOK
	}
	if o.StatusNG != "" {
		c.StatusNG = o.StatusNG
	}
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
