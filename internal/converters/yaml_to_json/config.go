package yaml_to_json

import (
	"fmt"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/common/validation"
)

type YamlToJsonConfig struct {
	IndentSize int    `json:"indent_size" validate:"required"`
	IndentType string `json:"indent_type" validate:"required,oneof=space tab"`
}

func (y YamlToJsonConfig) Validate() error {
	err := validation.GetValidator().Struct(y)
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

// FormattingConfig interface implementation
func (y YamlToJsonConfig) GetIndentType() string {
	return y.IndentType
}

func (y YamlToJsonConfig) GetIndentSize() int {
	return y.IndentSize
}

func (y YamlToJsonConfig) GetIndentString() string {
	return data.GetIndentString(y.IndentType, y.IndentSize)
}
