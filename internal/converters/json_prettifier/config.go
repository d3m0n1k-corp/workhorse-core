package json_prettifier

import (
	"fmt"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/common/validation"
)

type JsonPrettifierConfig struct {
	IndentSize int    `json:"indent_size" validate:"required"`
	IndentType string `json:"indent_type" validate:"required,oneof=space tab"`
}

func (j JsonPrettifierConfig) Validate() error {
	err := validation.GetValidator().Struct(j)
	if err != nil {
		return err
	}

	if j.IndentType == "space" && j.IndentSize%2 != 0 {
		return fmt.Errorf("Indent size for space should be even")
	}

	if j.IndentType == "tab" && j.IndentSize != 1 {
		return fmt.Errorf("Indent size for tab should be 1")
	}

	return nil
}

// FormattingConfig interface implementation
func (j JsonPrettifierConfig) GetIndentType() string {
	return j.IndentType
}

func (j JsonPrettifierConfig) GetIndentSize() int {
	return j.IndentSize
}

func (j JsonPrettifierConfig) GetIndentString() string {
	return data.GetIndentString(j.IndentType, j.IndentSize)
}
