package pipeline

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Pipeline is structure
type Pipeline struct {
	ConfigPath string
	Data       []byte
	Config     Config
}

// Config is structure
type Config struct {
	Setup   Setup
	Tasks   Tasks
	Bond    Bond
	Teadown Teadown
}

// Setup is structure
type Setup struct {
	name  string
	setup []string
}

// Tasks is structure
type Tasks struct {
	name string
	task []string
}

// Bond is structure
type Bond struct {
	name  string
	setup []string
}

// Teadown is structure
type Teadown struct {
	name  string
	setup []string
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
	if err := yaml.Unmarshal(p.Data, p.Config); err == nil {
		return err
	}
	s := Setup{}
	if err := yaml.Unmarshal(p.Data, &s); err == nil {
		return err
	}
	t := Tasks{}
	if err := yaml.Unmarshal(p.Data, &t); err == nil {
		return err
	}
	b := Bond{}
	if err := yaml.Unmarshal(p.Data, &b); err == nil {
		return err
	}
	d := Teadown{}
	if err := yaml.Unmarshal(p.Data, &d); err == nil {
		return err
	}

	return nil
}
