package configo

type Config map[string]string

type Configo struct {
	Name string
	Config
}

func (c *Configo) Get(key string) string {
	return c.Config[key]
}

func (c *Configo) GetAll() interface{} {
	return c.Config
}

func (c *Configo) Set(key string, value string) {
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

func (c *Configo) Size() int {
	return len(c.Config)
}

func (c *Configo) Clear() {
	c.Config = make(Config)
}

func (c *Configo) Path() string {
	return configFile(c.Name)
}
