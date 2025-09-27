package converters

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) Validate() error {
	m.Called()
	return nil
}

type TestConfig struct {
	Action string
}

func (t TestConfig) Validate() error {
	if t.Action != "success" {
		return fmt.Errorf("Invalid config")
	}
	return nil
}

func TestRegister_whenNewRegistration_returnNil(t *testing.T) {
	originalSize := GetRegistrySize()
	defer func() {
		// Clean up by removing the test registration
		delete(registry, "test")
	}()

	reg := Registration{
		Name: "test",
	}
	val := Register(&reg)
	require.Nil(t, val)
	require.Equal(t, originalSize+1, GetRegistrySize())
}

func TestRegister_whenAlreadyRegistered_panic(t *testing.T) {
	defer func() {
		// Clean up
		delete(registry, "test_panic")
	}()

	reg := Registration{
		Name: "test_panic",
	}
	// Register once
	registry[reg.Name] = &reg

	// Try to register again - should panic
	require.Panics(t, func() { Register(&reg) })
}

func TestGetRegistration_whenFound_returnRegistration(t *testing.T) {
	defer func() {
		delete(registry, "test_get")
	}()

	reg := Registration{
		Name: "test_get",
	}
	registry[reg.Name] = &reg
	val, err := GetRegistration(reg.Name)
	require.Nil(t, err)
	require.Equal(t, reg, *val)
}
func TestGetRegistration_whenNotFound_returnNil(t *testing.T) {
	_, err := GetRegistration("test")
	require.NotNil(t, err)
}

func TestNewConverter_whenRegistrationNotFound_returnError(t *testing.T) {
	_, err := NewConverter("test", "{}")
	require.Error(t, err)
}

func TestNewConverter_whenInvalidConfigJson_returnError(t *testing.T) {
	defer func() {
		delete(registry, "test_invalid_json")
	}()

	reg := Registration{
		Name:   "test_invalid_json",
		Config: reflect.TypeOf(MockConfig{}),
	}
	registry[reg.Name] = &reg
	_, err := NewConverter(reg.Name, "{")
	require.Error(t, err)
}

func TestNewConverter_whenInvalidConfigType_returnError(t *testing.T) {
	defer func() {
		delete(registry, "test_invalid_type")
	}()

	reg := Registration{
		Name:   "test_invalid_type",
		Config: reflect.TypeOf(struct{ Name string }{}),
	}
	registry[reg.Name] = &reg
	_, err := NewConverter(reg.Name, "{}")
	require.Error(t, err, "Invalid config type")
}

func TestNewConverter_whenInvalidConfigValues_returnError(t *testing.T) {
	defer func() {
		delete(registry, "test_invalid_values")
	}()

	reg := Registration{
		Name:   "test_invalid_values",
		Config: reflect.TypeOf(TestConfig{}),
	}
	registry[reg.Name] = &reg
	_, err := NewConverter(reg.Name, "{}")
	require.Error(t, err, "Invalid config")
}

func TestNewConverter_whenValidConfig_returnConverter(t *testing.T) {
	defer func() {
		delete(registry, "test_valid")
	}()

	reg := Registration{
		Name:   "test_valid",
		Config: reflect.TypeOf(TestConfig{}),
		Constructor: func(config BaseConfig) BaseConverter {
			return nil
		},
	}
	registry[reg.Name] = &reg
	config := TestConfig{Action: "success"}
	configStr, _ := json.Marshal(config)
	_, err := NewConverter(reg.Name, string(configStr))
	require.NoError(t, err)
}

func TestListConverters(t *testing.T) {
	// Save original registry state
	originalKeys := GetRegistryKeys()
	originalRegs := make(map[string]*Registration)
	for _, key := range originalKeys {
		originalRegs[key] = registry[key]
	}

	// Clean registry for test
	ResetRegistry()
	defer func() {
		// Restore original state
		ResetRegistry()
		for key, reg := range originalRegs {
			registry[key] = reg
		}
	}()

	registry["test_list"] = &Registration{Name: "test_list"}

	converters := ListConverters()
	require.Len(t, converters, 1)
	require.Equal(t, "test_list", converters[0].Name)
}
