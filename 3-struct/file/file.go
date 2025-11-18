package file

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

func IsJSONFile(fileName string) bool {
	ext := strings.ToLower(strings.TrimSpace(filepath.Ext(fileName)))
	return ext == ".json"
}
