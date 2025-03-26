package base

import(
	"encoding/base64"
	"encoding/base32"
	"fmt"
)

type DataBase struct{
	Data string
}

func (s *DataBase) InputType64() string{
	encodedValue := base64.StdEncoding.EncodeToString([]byte(s.Data))
	return encodedValue
}

func (s *DataBase) OutputType64() string {
	decodedBytes, err := base64.StdEncoding.DecodeString(s.Data)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return "" // Return empty string on failure
	}
	return string(decodedBytes) // Convert []byte to string
}

func (s *DataBase) InputType32() string{
	encodedValue := base32.StdEncoding.EncodeToString([]byte(s.Data))
	return encodedValue
}

func (s *DataBase) OutputType32() string {
	decodedBytes, err := base32.StdEncoding.DecodeString(s.Data)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return "" // Return empty string on failure
	}
	return string(decodedBytes)
}

