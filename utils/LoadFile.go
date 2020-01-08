package utils

import (
	"io/ioutil"
)

// LoadFile - Utility which can read the file and return its contents
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
