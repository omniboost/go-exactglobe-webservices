package utils

import (
	"encoding/json"
	"strings"
)

type JsonTester struct {
	json.RawMessage
}

func (j JsonTester) String() string {
	return string(j.RawMessage)
}

func (j JsonTester) IsObject() bool {
	return strings.HasPrefix(string(j.RawMessage), "{")
}

func (j JsonTester) IsArray() bool {
	return strings.HasPrefix(string(j.RawMessage), "[")
}
