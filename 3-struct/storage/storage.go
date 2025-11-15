package storage

import (
	"3-struct/bins"
	"3-struct/files"
	"encoding/json"
	"os"
)

func BinListToJSON(binList []bins.Bin, filename string) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadJSONFile(filePath string) ([]bins.Bin, error) {
	data, err := files.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var binList []bins.Bin
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, err
	}
	return binList, nil
}
