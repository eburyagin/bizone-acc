package ds

import (
	"os"
	"testing"

	cfg "github.com/eburyagin/bizone-acc/internal/cfg"
)

func TestConnection(t *testing.T) {
	cf := "../../" + os.Getenv("BZ_TEST_CONFIG_FILE")
	config, err := cfg.Load(cf)
	ds, err := NewConnect(config)
	if err != nil {
		t.Errorf("Connect(%q) = %v, error: %s", config, ds, err)
	}
}

func TestDisconnection(t *testing.T) {
	cf := "../../" + os.Getenv("BZ_TEST_CONFIG_FILE")
	config, _ := cfg.Load(cf)
	ds, _ := NewConnect(config)
	err := Disconnect(ds)
	if err != nil {
		t.Errorf("Disconnect(%q), error: %s", ds, err)
	}
}
