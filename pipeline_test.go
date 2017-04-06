package pipeline

import (
	"testing"

	"github.com/linyows/pipeline"
)

func TestNew(t *testing.T) {
	p := pipeline.New()
	expected := ".pipeline.yml"

	if p.ConfigPath != expected {
		t.Errorf("expected %s to eq %s", p.ConfigPath, expected)
	}

	if p.Data != nil {
		t.Errorf("expected %s to eq %s", p.Data, nil)
	}
}

func TestLoadConfig(t *testing.T) {
	p := pipeline.New()
	p.ConfigPath = "testdata/.pipeline.yml"
	p.LoadConfig()
	//expected := pipeline.Setup{}
	//if p.Config.Setup == expected {
	//	t.Errorf("%+v", p.Config)
	//}
}
