package performance_test

import (
	"strings"
	"testing"
	"workhorse-core/app"
	"workhorse-core/internal/chain"
	"workhorse-core/internal/common/data"
	"workhorse-core/internal/converters"
)

func BenchmarkCurrentStringChain(b *testing.B) {
	// Simulate a complex JSON object
	largeJSON := generateLargeJSON(1000) // 1000 nested objects

	chainLinks := []chain.ConverterChainLink{
		{Name: "json_prettifier", ConfigJSON: `{"indent_type":"space","indent_size":2}`},
		{Name: "json_to_yaml", ConfigJSON: `{}`},
		{Name: "yaml_to_json", ConfigJSON: `{"indent_type":"space","indent_size":2}`},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := app.ExecuteChain(chainLinks, largeJSON)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCurrentStringSingle(b *testing.B) {
	largeJSON := generateLargeJSON(1000)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := app.ExecuteConverter("json_prettifier", largeJSON, `{"indent_type":"space","indent_size":2}`)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkStructuredDataProcessing benchmarks the potential of structured data processing
func BenchmarkStructuredDataProcessing(b *testing.B) {
	largeJSON := generateLargeJSON(1000)

	// Create converters for manual structured testing
	prettifier, _ := converters.NewConverter("json_prettifier", `{"indent_type":"space","indent_size":2}`)
	jsonToYaml, _ := converters.NewConverter("json_to_yaml", `{}`)
	yamlToJson, _ := converters.NewConverter("yaml_to_json", `{"indent_type":"space","indent_size":2}`)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// Parse once
		structured, err := data.NewFromJSON(largeJSON)
		if err != nil {
			b.Fatal(err)
		}

		// Process with structured converters
		if prettifierStruct, ok := prettifier.(converters.StructuredConverter); ok && prettifierStruct.SupportsStructuredInput() {
			structured, err = prettifierStruct.ApplyStructured(structured)
			if err != nil {
				b.Fatal(err)
			}
		}

		if jsonToYamlStruct, ok := jsonToYaml.(converters.StructuredConverter); ok && jsonToYamlStruct.SupportsStructuredInput() {
			structured, err = jsonToYamlStruct.ApplyStructured(structured)
			if err != nil {
				b.Fatal(err)
			}
		}

		if yamlToJsonStruct, ok := yamlToJson.(converters.StructuredConverter); ok && yamlToJsonStruct.SupportsStructuredInput() {
			structured, err = yamlToJsonStruct.ApplyStructured(structured)
			if err != nil {
				b.Fatal(err)
			}
		}

		// Serialize only once at the end
		_, err = structured.ToJSON("  ")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func generateLargeJSON(size int) string {
	var builder strings.Builder
	builder.WriteString(`{"data":[`)

	for i := 0; i < size; i++ {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(`{"id":`)
		builder.WriteString(string(rune('0' + i%10)))
		builder.WriteString(`,"name":"item_`)
		builder.WriteString(string(rune('0' + i%10)))
		builder.WriteString(`","nested":{"value":`)
		builder.WriteString(string(rune('0' + i%10)))
		builder.WriteString(`}}`)
	}

	builder.WriteString(`]}`)
	return builder.String()
}
