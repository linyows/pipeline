package pipeline

import (
	"testing"
)

func TestNewPipeline(t *testing.T) {
	p := NewPipeline()
	expected := ".pipeline.yml"

	if p.ConfigPath != expected {
		t.Errorf("expected %s to eq %s", p.ConfigPath, expected)
	}

	if p.Data != nil {
		t.Errorf("expected %s to eq nil", p.Data)
	}
}

func TestLoadConfig(t *testing.T) {
	p := NewPipeline()
	p.ConfigPath = "testdata/.pipeline.toml"
	p.LoadConfig()
	if p.Config.Setup.name != "" {
		t.Errorf("%+v", p.Config.Setup.name)
	}
}
