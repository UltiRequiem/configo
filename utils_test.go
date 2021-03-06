package configo

import (
	"fmt"
	"testing"
)

func TestConfigFile(t *testing.T) {
	configurationFile := fmt.Sprintf("%s/example.json", configDir())
	toTestConfigurationFile := configFile("example")

	if !(configurationFile == toTestConfigurationFile) {
		t.Fatalf("Expected %s but got %s", configurationFile, toTestConfigurationFile)
	}
}
