package chain

import (
	"fmt"
	"testing"
	"workhorse-core/internal/converters"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockConverter1 struct {
	mock.Mock
}

func (m *MockConverter1) Apply(input any) (any, error) {
	m.Called()
	if input == nil {
		return nil, fmt.Errorf("error")
	}
	return nil, nil
}

func (m *MockConverter1) InputType() string {
	return "type1"
}

func (m *MockConverter1) OutputType() string {
	return "type2"
}

type MockConverter2 struct {
	mock.Mock
}

func (m *MockConverter2) Apply(input any) (any, error) {
	m.Called()
	if input == nil {
		return nil, fmt.Errorf("error")
	}
	return nil, nil
}

func (m *MockConverter2) InputType() string {
	return "type2"
}

func (m *MockConverter2) OutputType() string {
	return "type3"
}

func TestNewConverterListFromJSON_whenNoChainLinks_returnError(t *testing.T) {
	cl, err := NewConverterListFromJSON([]ConverterChainLink{})
	require.NotNil(t, err)
	require.Nil(t, cl)
}

func TestNewConverterListFromJSON_whenProperChainLinks_returnConverterList(t *testing.T) {
	mockableNewConverterFunc = func(name string, _ string) (converters.BaseConverter, error) {
		if name == "name1" {
			return &MockConverter1{}, nil
		}
		if name == "name2" {
			return &MockConverter2{}, nil
		}
		return nil, fmt.Errorf("error")
	}

	link1 := ConverterChainLink{
		Name:       "name1",
		ConfigJSON: "config1",
	}
	link2 := ConverterChainLink{
		Name:       "name2",
		ConfigJSON: "config2",
	}
	chainLinks := []ConverterChainLink{link1, link2}

	cl, err := NewConverterListFromJSON(chainLinks)
	require.Nil(t, err)
	require.NotNil(t, cl)
}

func TestNewConverterListFromJSON_whenInvalidChainLinks_returnError(t *testing.T) {
	mockableNewConverterFunc = func(name string, _ string) (converters.BaseConverter, error) {
		if name == "name1" {
			return &MockConverter1{}, nil
		}
		if name == "name2" {
			return &MockConverter2{}, nil
		}
		return nil, fmt.Errorf("error")
	}

	link1 := ConverterChainLink{
		Name:       "error",
		ConfigJSON: "config1",
	}
	chainLinks := []ConverterChainLink{link1}

	cl, err := NewConverterListFromJSON(chainLinks)
	require.NotNil(t, err)
	require.Nil(t, cl)
}

func TestNewConverterListFromJSON_whenUnchainable_returnError(t *testing.T) {
	mockableNewConverterFunc = func(name string, _ string) (converters.BaseConverter, error) {
		if name == "name1" {
			return &MockConverter1{}, nil
		}
		if name == "name2" {
			return &MockConverter2{}, nil
		}
		return nil, fmt.Errorf("error")
	}

	link1 := ConverterChainLink{
		Name:       "name1",
		ConfigJSON: "config1",
	}
	link2 := ConverterChainLink{
		Name:       "name2",
		ConfigJSON: "config2",
	}
	chainLinks := []ConverterChainLink{link2, link1}

	cl, err := NewConverterListFromJSON(chainLinks)
	require.NotNil(t, err)
	require.Nil(t, cl)
}
