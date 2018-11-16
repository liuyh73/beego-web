package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, _ := os.Open(path)
	defer file.Close()
	return ioutil.ReadAll(file)
}
