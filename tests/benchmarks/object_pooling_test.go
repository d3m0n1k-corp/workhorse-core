package performance_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"workhorse-core/internal/common/pools"
)

// BenchmarkJSONMarshalWithoutPooling benchmarks JSON marshaling without object pooling
func BenchmarkJSONMarshalWithoutPooling(b *testing.B) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
		"data": map[string]interface{}{
			"nested": "value",
			"array":  []int{1, 2, 3, 4, 5},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetIndent("", "  ")
		encoder.SetEscapeHTML(false)
		_ = encoder.Encode(data)
		_ = buf.String()
	}
}

// BenchmarkJSONMarshalWithPooling benchmarks JSON marshaling with object pooling
func BenchmarkJSONMarshalWithPooling(b *testing.B) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
		"data": map[string]interface{}{
			"nested": "value",
			"array":  []int{1, 2, 3, 4, 5},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result string
		_ = pools.GlobalPool.WithBuffer(func(buf *bytes.Buffer) error {
			encoder := json.NewEncoder(buf)
			encoder.SetIndent("", "  ")
			encoder.SetEscapeHTML(false)
			if err := encoder.Encode(data); err != nil {
				return err
			}
			result = buf.String()
			return nil
		})
		_ = result
	}
}

// BenchmarkJSONProcessingChain benchmarks a chain of JSON operations
func BenchmarkJSONProcessingChain(b *testing.B) {
	input := `{"test": true, "number": 42, "nested": {"array": [1,2,3,4,5]}}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulate chain processing without pooling
		var data interface{}
		json.Unmarshal([]byte(input), &data)

		// Marshal prettified
		var buf1 bytes.Buffer
		enc1 := json.NewEncoder(&buf1)
		enc1.SetIndent("", "  ")
		enc1.Encode(data)
		pretty := buf1.String()

		// Marshal compact
		compact, _ := json.Marshal(data)
		_ = string(compact)
		_ = pretty
	}
}

// BenchmarkJSONProcessingChainWithPooling benchmarks with object pooling
func BenchmarkJSONProcessingChainWithPooling(b *testing.B) {
	input := `{"test": true, "number": 42, "nested": {"array": [1,2,3,4,5]}}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulate chain processing with pooling
		var data interface{}
		json.Unmarshal([]byte(input), &data)

		var pretty, compact string
		// Marshal prettified with pooling
		pools.GlobalPool.WithBuffer(func(buf *bytes.Buffer) error {
			enc := json.NewEncoder(buf)
			enc.SetIndent("", "  ")
			enc.Encode(data)
			pretty = buf.String()
			return nil
		})

		// Marshal compact with pooling
		pools.GlobalPool.WithBuffer(func(buf *bytes.Buffer) error {
			enc := json.NewEncoder(buf)
			enc.Encode(data)
			compact = buf.String()
			return nil
		})

		_ = pretty
		_ = compact
	}
}
