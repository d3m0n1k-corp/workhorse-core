package app

import (
	"fmt"
	"reflect"
	"workhorse-core/internal/chain"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/converters"

	"github.com/sirupsen/logrus"
)

type ChainResponse struct {
	Output any    `json:"output"`
	Error  string `json:"error"`
}

func ExecuteChain(chainLinks []chain.ConverterChainLink, input string) ([]*ChainResponse, error) {
	cl, err := chain.NewConverterListFromJSON(chainLinks)
	if err != nil {
		logrus.Errorf("Error creating converter list: %v", err)
		return nil, err
	}
	logrus.Trace("Converter list created successfully")
	err = cl.Validate()
	if err != nil {
		logrus.Errorf("Error validating converter list: %v", err)
		return nil, err
	}
	logrus.Trace("Converter list validated successfully")

	// Enable structured chain execution for supported converters
	if canUseStructuredChain(cl) {
		logrus.Trace("Using optimized structured chain execution")
		return executeStructuredChain(cl, input)
	}

	// Use string-based chain execution for compatibility
	logrus.Trace("Using string-based chain execution")
	return executeStringChain(cl, input)
}

// canUseStructuredChain checks if all converters in the chain support structured execution
func canUseStructuredChain(cl *chain.ConverterList) bool {
	head := cl.Head()
	for head != nil {
		if structuredConv, ok := head.Value.(converters.StructuredConverter); !ok || !structuredConv.SupportsStructuredInput() {
			return false
		}
		head = head.Next
	}
	return true
}

// executeStructuredChain runs the chain using structured data to avoid repeated serialization
func executeStructuredChain(cl *chain.ConverterList, input string) ([]*ChainResponse, error) {
	head := cl.Head()
	response := make([]*ChainResponse, 0, cl.Length())

	// Parse input once based on the first converter's expected input type
	var currentData *data.IntermediateData
	var err error

	firstConverter := head.Value
	switch firstConverter.InputType() {
	case "json":
		currentData, err = data.NewFromJSON(input)
		if err != nil {
			// If structured parsing fails, fall back to string-based execution
			logrus.Trace("Structured parsing failed, falling back to string-based execution")
			return executeStringChain(cl, input)
		}
	case "yaml":
		currentData, err = data.NewFromYAML(input)
		if err != nil {
			// If structured parsing fails, fall back to string-based execution
			logrus.Trace("Structured parsing failed, falling back to string-based execution")
			return executeStringChain(cl, input)
		}
	default:
		// Unsupported input type, fall back to string-based
		return executeStringChain(cl, input)
	}

	// Execute chain using structured data with context-aware serialization
	for i := 0; head != nil; i++ {
		logrus.Tracef("Executing structured converter %d: %s", i, reflect.TypeOf(head.Value))

		structuredConv := head.Value.(converters.StructuredConverter)
		result, err := structuredConv.ApplyStructured(currentData)

		// Context-aware serialization: preserve converter-specific formatting
		var outputStr string
		if err == nil {
			outputStr, err = serializeWithContext(result, head.Value)
		}

		resp := &ChainResponse{
			Output: outputStr,
			Error:  fmt.Sprintf("%v", err),
		}
		response = append(response, resp)

		if err != nil {
			break
		}

		currentData = result
		head = head.Next
	}

	return response, nil
}

// serializeForConverter converts structured data to string format expected by converter
func serializeForConverter(intermediate *data.IntermediateData, inputType string) (string, error) {
	switch inputType {
	case "json":
		return intermediate.ToJSON("")
	case "yaml":
		return intermediate.ToYAML()
	case "json_stringified":
		return intermediate.ToJSONStringified()
	default:
		return "", fmt.Errorf("unsupported input type: %s", inputType)
	}
}

// executeStringChain runs the original string-based chain execution
func executeStringChain(cl *chain.ConverterList, input string) ([]*ChainResponse, error) {
	head := cl.Head()
	response := make([]*ChainResponse, 0, cl.Length())
	chain_in := input

	for i := 0; head != nil; i++ {
		logrus.Tracef("Executing string converter %d: %s", i, reflect.TypeOf(head.Value))
		out, err := head.Value.Apply(chain_in)
		resp := &ChainResponse{
			Output: out,
			Error:  fmt.Sprintf("%v", err),
		}

		response = append(response, resp)

		if err != nil {
			break
		}
		chain_in = out.(string)
		head = head.Next
	}

	return response, nil
}

// serializeWithContext converts structured data with converter-specific formatting
func serializeWithContext(intermediate *data.IntermediateData, converter converters.BaseConverter) (string, error) {
	switch converter.OutputType() {
	case "json":
		// Check if this converter has formatting configuration
		if configurable, ok := converter.(converters.ConfigurableConverter); ok {
			config := configurable.GetFormattingConfig()
			return intermediate.ToJSON(config.GetIndentString())
		}
		return intermediate.ToJSON("")
	case "yaml":
		return intermediate.ToYAML()
	case "json_stringified":
		// For JSON stringify converters, the data is already in string form
		// The ProcessData already did the stringify operation, so just return it as-is
		if result, ok := intermediate.Data.(string); ok {
			return result, nil
		}
		// Fallback: use the ToJSONStringified method if data isn't already stringified
		return intermediate.ToJSONStringified()
	default:
		return "", fmt.Errorf("unsupported output type: %s", converter.OutputType())
	}
}
