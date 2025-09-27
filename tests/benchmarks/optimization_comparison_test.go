package performance_test

import (
	"encoding/json"
	"strings"
	"testing"
	"workhorse-core/app"
	"workhorse-core/internal/chain"
	"workhorse-core/internal/common/data"
)

// BenchmarkChainExecutionOptimizationComparison compares the performance of optimized vs baseline chain execution
func BenchmarkChainExecutionOptimizationComparison(b *testing.B) {
	// Test data: JSON -> JSON Prettifier -> JSON Stringify -> JSON to YAML chain
	inputJSON := `{"users":[{"name":"Alice","age":30,"preferences":{"theme":"dark","notifications":true}},{"name":"Bob","age":25,"preferences":{"theme":"light","notifications":false}}],"metadata":{"version":"1.0","timestamp":"2024-01-01T00:00:00Z"}}`

	// Setup chain configuration - a valid conversion chain
	chainLinks := []chain.ConverterChainLink{
		{Name: "json_prettifier", ConfigJSON: `{"indent_type":"space","indent_size":2}`},
		{Name: "json_to_yaml", ConfigJSON: `{}`},
		{Name: "yaml_to_json", ConfigJSON: `{"indent_type":"space","indent_size":2}`},
	}

	b.ResetTimer()

	b.Run("OptimizedChainExecution", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Use the optimized chain execution
			_, err := app.ExecuteChain(chainLinks, inputJSON)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("IndividualConverterCalls", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Simulate individual converter calls (baseline approach)
			currentData := inputJSON

			// Json Prettifier
			result1, err := app.ExecuteConverter("json_prettifier", currentData, `{"indent_type":"space","indent_size":2}`)
			if err != nil {
				b.Fatal(err)
			}
			currentData = result1.(string)

			// Json to YAML
			result2, err := app.ExecuteConverter("json_to_yaml", currentData, `{}`)
			if err != nil {
				b.Fatal(err)
			}
			currentData = result2.(string)

			// YAML to Json
			_, err = app.ExecuteConverter("yaml_to_json", currentData, `{"indent_type":"space","indent_size":2}`)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkMemoryOptimization compares memory allocation patterns
func BenchmarkMemoryOptimization(b *testing.B) {
	inputData := make(map[string]interface{})
	for i := 0; i < 100; i++ {
		inputData[strings.Repeat("key", i)] = strings.Repeat("value", i*10)
	}

	b.Run("WithIntermediateDataStructures", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Create intermediate data structure (optimized path)
			intermediate := data.NewFromData(inputData, "json")

			// Simulate chain processing with structured data
			for j := 0; j < 3; j++ {
				jsonStr, _ := intermediate.ToJSON("")
				intermediate, _ = data.NewFromJSON(jsonStr)
			}
		}
	})

	b.Run("WithStringConversions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Simulate repeated string serialization/deserialization
			intermediate := data.NewFromData(inputData, "json")
			currentStr, _ := intermediate.ToJSON("")

			for j := 0; j < 3; j++ {
				var tempData interface{}
				json.Unmarshal([]byte(currentStr), &tempData)
				tempBytes, _ := json.Marshal(tempData)
				currentStr = string(tempBytes)
			}
		}
	})
}
