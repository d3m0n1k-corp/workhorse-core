package base_test

import (
	"testing"
	base "workhorse-core/internal/converters/base_converter"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/suite"
	// "github.com/your/project/yourpackage"
)

func TestNewEncoding64(t *testing.T) {
	// Test the NewEncoding64 function
	val := base.DataBase{Data: "Hello, Go!"}
	expectedOutput := "SGVsbG8sIEdvIQ=="

	encoded := val.InputType64()

	if encoded != expectedOutput {
		t.Fatal("NewEncoding returned nil")
	}
}

func TestNewDecoding64(t *testing.T){
	// Test the NewDecoding64 function
	val := base.DataBase{Data: "SGVsbG8sIEdvIQ=="}
	expectedOutput := "Hello, Go!"

	encoded := val.OutputType64()

	if encoded != expectedOutput{
		t.Fatal("NewDecoding returned nil")
	}
}