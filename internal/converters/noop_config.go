package converters

// NoOpConfig is a shared empty configuration struct for converters that don't require configuration
type NoOpConfig struct{}

// Validate implements BaseConfig interface with no-op validation
func (c NoOpConfig) Validate() error {
	return nil
}
