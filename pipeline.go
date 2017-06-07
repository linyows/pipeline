package pipeline

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// Pipeline is structure
type Pipeline struct {
	ConfigPath string
	Data       []byte
	Config     Config
	Lines      []*Tasks
}

// Config is structure
type Config struct {
}

// Tasks is structure
type Task struct {
	name    string
	command string
}

// NewPipeline for pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{
		ConfigPath: ".pipeline.yml",
	}
}

// Run invokes the CLI with the given arguments.
func (p *Pipeline) Run(args []string) int {
	if err := p.LoadConfig(); err != nil {
		return 1
	}
	return 0
}

// LoadConfig loads a config file
func (p *Pipeline) LoadConfig() error {
	data, err := ioutil.ReadFile(p.ConfigPath)
	if err != nil {
		return err
	}

	p.Data = data
	var conf Config

	if _, err := yaml.Unmarshal(data, p); err != nil {
		return err
	}

	return nil
}
