package packlit

import "fmt"

// Generic flag types for extensibility
type GenericKeyValueOption struct {
	Key   string
	Value string
}

func (g GenericKeyValueOption) Validate() error { return nil }
func (g GenericKeyValueOption) Parse() string   { return fmt.Sprintf("%s=%s", g.Key, g.Value) }

type GenericRawFlag string

func (g GenericRawFlag) Validate() error { return nil }
func (g GenericRawFlag) Parse() string   { return string(g) }

type GenericKeyValueFlag struct {
	Key   string
	Value string
}

func (g GenericKeyValueFlag) Validate() error {
	if g.Key == "" {
		return fmt.Errorf("flag key cannot be empty")
	}
	return nil
}

func (g GenericKeyValueFlag) Parse() string { return fmt.Sprintf("--%s=%s", g.Key, g.Value) }

type GenericBoolFlag string

func (g GenericBoolFlag) Validate() error {
	if string(g) == "" {
		return fmt.Errorf("flag name cannot be empty")
	}
	return nil
}

func (g GenericBoolFlag) Parse() string { return "--" + string(g) }
