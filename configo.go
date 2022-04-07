package configo

import (
	"encoding/json"
	"io/ioutil"
)

type Config map[string]interface{}

type Configo struct {
	Name string
	Config
}

// Returns the `value` on the `key` passed.
func (c *Configo) Get(key string) interface{} {
	return c.Config[key]
}

// Returns a map with all the config.
func (c *Configo) GetAll() interface{} {
	return c.Config
}

// Set the `value` in the `key` passed.
func (c *Configo) Set(key string, value interface{}) {
	c.Config[key] = value
}

func (c *Configo) SetAll(config Config) {
	c.Config = config
}

// It tells you if it has the received `key`.
func (c *Configo) Has(key string) bool {
	if _, ok := c.Config[key]; ok {
		return true
	}

	return false
}

// Deletes the `key` passed.
func (c *Configo) Delete(key string) {
	delete(c.Config, key)
}

// Returns the quantity of the config properties.
func (c *Configo) Size() int {
	return len(c.Config)
}

// Removes all the values of the config.
func (c *Configo) Clear() {
	c.Config = make(Config)
}

// Gives you the path were the config is stored.
func (c *Configo) Path() string {
	return configFile(c.Name)
}

// Saves the config for persistence.
func (c *Configo) Save() error {
	data, errorParsingJSON := json.MarshalIndent(c.Config, "", "  ")

	if errorParsingJSON != nil {
		return errorParsingJSON
	}

	errorWrittingFile := ioutil.WriteFile(c.Path(), data, 0644)

	if errorWrittingFile != nil {
		return errorWrittingFile
	}

	return nil
}
