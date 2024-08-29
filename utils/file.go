package utils

import (
	"os"
)

type File struct {}

func (j File) Save(path string, data string) error {
	dataInByte := []byte(data)
	err := os.WriteFile(path, dataInByte, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (j File) Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}