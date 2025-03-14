package json_prettifier

import "testing"

func TestInputType_returnType(t *testing.T) {
	j := JsonPrettifier{}
	if j.InputType() != "json" {
		t.Error("Input type should be json")
	}
}

func TestOutputType_returnType(t *testing.T) {
	j := JsonPrettifier{}
	if j.OutputType() != "json" {
		t.Error("Output type should be json")
	}
}

func TestApply_whenValidSpaces_returnPrettyJson(t *testing.T) {
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 4, IndentType: "space"}}
	input := []byte(`{"a":1,"b":2}`)
	output, err := j.Apply(input)
	if err != nil {
		t.Errorf("Error while applying json prettifier: %v", err)
	}
	expected := "{\n    \"a\": 1,\n    \"b\": 2\n}"
	if string(output.([]byte)) != expected {
		t.Errorf("Expected %s, got %s", expected, string(output.([]byte)))
	}
}

func TestApply_whenInvalidInput_returnError(t *testing.T) {
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 4, IndentType: "space"}}
	input := []byte(`{"a":1,"b":2`)
	_, err := j.Apply(input)
	if err == nil {
		t.Error("Error expected while applying json prettifier with invalid input")
	}
}

func TestApply_whenValidTabs_returnPrettyJson(t *testing.T) {
	j := JsonPrettifier{config: JsonPrettifierConfig{IndentSize: 1, IndentType: "tab"}}
	input := []byte(`{"a":1,"b":2}`)
	output, err := j.Apply(input)
	if err != nil {
		t.Errorf("Error while applying json prettifier: %v", err)
	}
	expected := "{\n\t\"a\": 1,\n\t\"b\": 2\n}"
	if string(output.([]byte)) != expected {
		t.Errorf("Expected %s, got %s", expected, string(output.([]byte)))
	}
}
