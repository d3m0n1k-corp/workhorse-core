package base_64

import (
	"bytes"
	"testing"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/suite"
	// "github.com/your/project/yourpackage"
)

func TestNewEncoding(t *testing.T) {
	e := NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	if e == nil {
		t.Fatal("NewEncoding returned nil")
	}
}

func TestEncode(t *testing.T) {
	enc := StdEncoding
	input := []byte("hello")
	expected := "aGVsbG8="

	output := make([]byte, enc.EncodedLen(len(input)))
	enc.Encode(output, input)

	if string(output) != expected {
		t.Errorf("Encode failed: got %s, expected %s", string(output), expected)
	}
}

func TestEncodeToString(t *testing.T) {
	enc := StdEncoding
	input := []byte("hello")
	expected := "aGVsbG8="
	output := enc.EncodeToString(input)

	if output != expected {
		t.Errorf("EncodeToString failed: got %s, expected %s", output, expected)
	}
}

func TestDecode(t *testing.T) {
	enc := StdEncoding
	input := "aGVsbG8="
	expected := "hello"

	output := make([]byte, enc.DecodedLen(len(input)))
	n, err := enc.Decode(output, []byte(input))
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if string(output[:n]) != expected {
		t.Errorf("Decode failed: got %s, expected %s", string(output[:n]), expected)
	}
}

func TestDecodeString(t *testing.T) {
	enc := StdEncoding
	input := "aGVsbG8="
	expected := "hello"

	output, err := enc.DecodeString(input)
	if err != nil {
		t.Fatalf("DecodeString failed: %v", err)
	}

	if string(output) != expected {
		t.Errorf("DecodeString failed: got %s, expected %s", string(output), expected)
	}
}

func TestEncodingWithPadding(t *testing.T) {
	enc := StdEncoding.WithPadding(NoPadding)
	if enc.padChar != NoPadding {
		t.Errorf("WithPadding failed: expected NoPadding, got %v", enc.padChar)
	}
}

func TestEncodingStrictMode(t *testing.T) {
	enc := StdEncoding.Strict()
	if !enc.strict {
		t.Errorf("Strict mode not set")
	}
}

func TestNewEncoder(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(StdEncoding, &buf)
	_, err := e.Write([]byte("hello"))
	if err != nil {
		t.Fatalf("NewEncoder Write failed: %v", err)
	}
	e.Close()

	expected := "aGVsbG8="
	if buf.String() != expected {
		t.Errorf("NewEncoder failed: got %s, expected %s", buf.String(), expected)
	}
}
