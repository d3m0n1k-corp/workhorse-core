package converters

import "workhorse-core/internal/common/data"

// FormattingConfig interface for converters that need specific formatting
type FormattingConfig interface {
	GetIndentType() string
	GetIndentSize() int
	GetIndentString() string
}

// ConfigurableConverter interface for converters that expose their formatting config
type ConfigurableConverter interface {
	BaseConverter
	GetFormattingConfig() FormattingConfig
}

// StructuredConverter interface for high-performance converters
// All converters should implement this for optimal chain performance
type StructuredConverter interface {
	BaseConverter
	ApplyStructured(input *data.IntermediateData) (*data.IntermediateData, error)
	SupportsStructuredInput() bool
}

// DualModeConverter combines both string and structured processing
// This is what all converters should implement going forward
type DualModeConverter interface {
	StructuredConverter
	ConfigurableConverter

	// Core processing logic that both modes can use
	ProcessData(data any, config BaseConfig) (any, error)
}
