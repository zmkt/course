package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func StorageBinToFile(filename string, bins interface{}) error {

	data, err := json.Marshal(filename)
	if err != nil {
		fmt.Println("Ошибка")
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadFileDin(fileName string, bin interface{}) error {

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка")
		return err
	}

	return json.Unmarshal(data, bin)
}
