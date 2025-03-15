package converters

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegister_whenNewRegistration_returnNil(t *testing.T) {
	reg := Registration{
		Name: "test",
	}
	val := Register(reg)
	require.Nil(t, val)
}

func TestRegister_whenAlreadyRegistered_panic(t *testing.T) {
	reg := Registration{
		Name: "test",
	}
	registry[reg.Name] = reg
	require.Panics(t, func() { Register(reg) })
	delete(registry, reg.Name)
}

func TestGetRegistration_whenFound_returnRegistration(t *testing.T) {
	reg := Registration{
		Name: "test",
	}
	registry[reg.Name] = reg
	val, err := GetRegistration(reg.Name)
	require.Nil(t, err)
	require.Equal(t, reg, *val)
	delete(registry, reg.Name)
}
func TestGetRegistration_whenNotFound_returnNil(t *testing.T) {
	_, err := GetRegistration("test")
	require.NotNil(t, err)
}
