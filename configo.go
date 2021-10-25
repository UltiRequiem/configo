package configo

type Config map[interface{}]interface{}

type Configo struct {
	Name string
	Config
}

func (c *Configo) Get(key interface{}) interface{} {
	return c.Config[key]
}

func (c *Configo) GetAll() interface{} {
	return c.Config
}

func (c *Configo) Set(key interface{}, value interface{}) {
	c.Config[key] = value
}

func (c *Configo) SetAll(config Config) {
	c.Config = config
}

func (c *Configo) Has(key interface{}) bool {
	if _, ok := c.Config[key]; ok {
		return true
	}
	return false
}

func (c *Configo) Delete(key interface{}) {
	c.Config[key] = nil
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
