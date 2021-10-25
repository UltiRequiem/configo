// Dead Simple Configuration Management
package configo

func NewConfig(name string) (*Configo, error) {
	return &Configo{name, make(Config)}, nil
}
