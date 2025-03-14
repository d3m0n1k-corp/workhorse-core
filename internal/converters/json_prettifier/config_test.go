package json_prettifier

import (
	"encoding/json"
	"testing"
)

func TestJsonPrettyfierConfig_whenValid_returnNil(t *testing.T) {
	conf_json := `{"indent_size": 4, "indent_type": "space"}`
	var conf JsonPrettifierConfig
	err := json.Unmarshal([]byte(conf_json), &conf)
	if err != nil {
		t.Error("Error while unmarshalling JsonPrettifierConfig")
	}
	err = conf.Validate()
	if err != nil {
		t.Error("Error while validating JsonPrettifierConfig")
	}
}

func TestJsonPrettyfierConfig_whenInvalidIndentType_returnError(t *testing.T) {
	conf_json := `{"indent_size": 4, "indent_type": "none"}`
	var conf JsonPrettifierConfig
	err := json.Unmarshal([]byte(conf_json), &conf)
	if err != nil {
		t.Errorf("Error while unmarshalling JsonPrettifierConfig: %v", err)
	}
	err = conf.Validate()
	if err == nil {
		t.Errorf("Error expected while validating JsonPrettifierConfig with invalid indent type")
	}
}

func TestJsonPrettyfierConfig_whenOddSpaces_returnError(t *testing.T) {
	conf_json := `{"indent_size": 3, "indent_type": "space"}`
	var conf JsonPrettifierConfig
	err := json.Unmarshal([]byte(conf_json), &conf)
	if err != nil {
		t.Errorf("Error while unmarshalling JsonPrettifierConfig: %v", err)
	}

	err = conf.Validate()
	if err == nil {
		t.Error("Error expected while validating JsonPrettifierConfig with space indent type and indent size odd")
	}

}

func TestJsonPrettyfierConfig_whenTabIndentSizeNotOne_returnError(t *testing.T) {
	conf_json := `{"indent_size": 2, "indent_type": "tab"}`
	var conf JsonPrettifierConfig
	err := json.Unmarshal([]byte(conf_json), &conf)
	if err != nil {
		t.Errorf("Error while unmarshalling JsonPrettifierConfig: %v", err)
	}

	err = conf.Validate()
	if err == nil {
		t.Error("Error expected while validating JsonPrettifierConfig with tab indent type and indent size not 1")
	}
}
