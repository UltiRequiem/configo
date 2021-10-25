package configo

import (
	"fmt"
	"os"
)

func configFile(file string) string {
	return fmt.Sprintf("%s/%s.json", configDir(), file)
}

func configDir() string {
	return fmt.Sprintf("%s/configo", os.Getenv("XDG_CONFIG_HOME"))
}
