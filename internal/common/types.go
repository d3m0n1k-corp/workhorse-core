package common

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type DataType int

const (
	JSON DataType = iota
	XML
	YAML
	BASE64
	BYTES
	DOT_ENV
)

var dataMap = map[DataType]string{
	JSON:    "json",
	XML:     "xml",
	YAML:    "yaml",
	BASE64:  "base64",
	BYTES:   "bytes",
	DOT_ENV: "env",
}

func (d DataType) String() string {
	return dataMap[d]
}

func GetDataType(s string) (DataType, error) {
	for k, v := range dataMap {
		if v == s {
			return k, nil
		}
	}
	err := fmt.Errorf("Unknown data type %s", s)
	logrus.Error(err)
	return -1, err
}
