package validation

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	instance *validator.Validate
	once     sync.Once
)

// GetValidator returns a singleton validator instance
func GetValidator() *validator.Validate {
	once.Do(func() {
		instance = validator.New()
	})
	return instance
}
