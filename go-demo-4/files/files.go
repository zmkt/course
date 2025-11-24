package files

import (
	"fmt"
	"os"
)

type JsonDB struct {
	fileName string
}

func NewJsonDB(name string) *JsonDB {

	return &JsonDB{
		fileName: name,
	}

}

func (db *JsonDB) Read() ([]byte, error) {

	data, err := os.ReadFile(db.fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *JsonDB) Write(content []byte) {

	file, err := os.Create(db.fileName)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")
}
