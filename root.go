// Dead Simple Config Management, load and persist config without having to think about where and how.
package configo

import (
	"encoding/json"
	"os"
)

// Returns a new instance of Configo
func NewConfig(name string) (*Configo, error) {
	configFilePath := configFile(name)

	configo := &Configo{name, make(Config)}

	if _, file := os.Stat(configFilePath); !os.IsNotExist(file) {
		data, errorReadingFile := os.ReadFile(configFilePath)

		if errorReadingFile != nil {
			return nil, errorReadingFile
		}

		m := map[string]interface{}{}

		err := json.Unmarshal(data, &m)

		if err != nil {
			return nil, err
		}

		for k, v := range m {
			configo.Set(string(k), v)
		}
	}

	return configo, nil
}
