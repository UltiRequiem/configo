// Dead Simple Config Management, easily load and persist config without having to think about where and how.
package configo

func NewConfig(name string) (*Configo, error) {
	return &Configo{name, make(Config)}, nil
}
