package files

import (
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
