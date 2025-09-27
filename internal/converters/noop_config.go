package converters

// NoOpConfig is a shared empty configuration struct for converters that don't require configuration
type NoOpConfig struct{}

// Validate implements BaseConfig interface with no-op validation
func (c NoOpConfig) Validate() error {
	return nil
}

// NoOpFormattingConfig for converters that don't have specific formatting requirements
type NoOpFormattingConfig struct{}

// GetIndentType returns empty string for no formatting
func (c NoOpFormattingConfig) GetIndentType() string {
	return ""
}

// GetIndentSize returns 0 for no formatting
func (c NoOpFormattingConfig) GetIndentSize() int {
	return 0
}

// GetIndentString returns empty string for no formatting
func (c NoOpFormattingConfig) GetIndentString() string {
	return ""
}
