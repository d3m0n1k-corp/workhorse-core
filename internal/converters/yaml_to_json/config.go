package yaml_to_json

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var vd = validator.New()

type YamlToJsonConfig struct {
	IndentSize int    `json:"indent_size" validate:"required"`
	IndentType string `json:"indent_type" validate:"required,oneof=space tab"`
}

func (y *YamlToJsonConfig) Validate() error {
	err := vd.Struct(y)
	if err != nil {
		return err
	}

	if y.IndentType == "space" && y.IndentSize%2 != 0 {
		return fmt.Errorf("Indent size for space should be even")
	}

	if y.IndentType == "tab" && y.IndentSize != 1 {
		return fmt.Errorf("Indent size for tab should be 1")
	}

	return nil
}
