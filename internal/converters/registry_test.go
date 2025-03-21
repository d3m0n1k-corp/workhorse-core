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

func (t *TestConfig) Validate() error {
	if t.Action != "success" {
		return fmt.Errorf("Invalid config")
	}
	return nil
}

func TestRegister_whenNewRegistration_returnNil(t *testing.T) {
	reg := Registration{
		Name: "test",
	}
	val := Register(&reg)
	require.Nil(t, val)
}

func TestRegister_whenAlreadyRegistered_panic(t *testing.T) {
	reg := Registration{
		Name: "test",
	}
	registry[reg.Name] = &reg
	require.Panics(t, func() { Register(&reg) })
	delete(registry, reg.Name)
}

func TestGetRegistration_whenFound_returnRegistration(t *testing.T) {
	reg := Registration{
		Name: "test",
	}
	registry[reg.Name] = &reg
	val, err := GetRegistration(reg.Name)
	require.Nil(t, err)
	require.Equal(t, reg, *val)
	delete(registry, reg.Name)
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
	reg := Registration{
		Name:   "test",
		Config: reflect.TypeOf(MockConfig{}),
	}
	registry[reg.Name] = &reg
	_, err := NewConverter(reg.Name, "{")
	require.Error(t, err)
	delete(registry, reg.Name)
}

func TestNewConverter_whenInvalidConfigType_returnError(t *testing.T) {
	reg := Registration{
		Name:   "test",
		Config: reflect.TypeOf(struct{ Name string }{}),
	}
	registry[reg.Name] = &reg
	_, err := NewConverter(reg.Name, "{}")
	require.Error(t, err, "Invalid config type")
	delete(registry, reg.Name)
}

func TestNewConverter_whenInvalidConfigValues_returnError(t *testing.T) {
	reg := Registration{
		Name:   "test",
		Config: reflect.TypeOf(TestConfig{}),
	}
	registry[reg.Name] = &reg
	_, err := NewConverter(reg.Name, "{}")
	require.Error(t, err, "Invalid config")
	delete(registry, reg.Name)
}

func TestNewConverter_whenValidConfig_returnConverter(t *testing.T) {
	reg := Registration{
		Name:   "test",
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
	delete(registry, reg.Name)
}

func TestListConverters(t *testing.T) {
	registry = make(map[string]*Registration)
	registry["test"] = &Registration{Name: "test"}

	converters := ListConverters()
	require.Len(t, converters, 1)
	require.Equal(t, "test", converters[0].Name)
}
