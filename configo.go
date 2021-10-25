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

func (c *Configo) Get(key string) interface{} {
	return c.Config[key]
}

func (c *Configo) GetAll() interface{} {
	return c.Config
}

func (c *Configo) Set(key string, value interface{}) {
	c.Config[key] = value
}

func (c *Configo) SetAll(config Config) {
	c.Config = config
}

func (c *Configo) Has(key string) bool {
	if _, ok := c.Config[key]; ok {
		return true
	}

	return false
}

func (c *Configo) Delete(key string) {
	delete(c.Config, key)
}

func (c *Configo) Size() int {
	return len(c.Config)
}

func (c *Configo) Clear() {
	c.Config = make(Config)
}

func (c *Configo) Path() string {
	return configFile(c.Name)
}

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
