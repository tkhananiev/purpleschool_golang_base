package files

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func IsJSON(filename string) bool {
	return strings.ToLower(filepath.Ext(filename)) == ".json"
}
