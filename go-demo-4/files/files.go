package files

import (
	"go-demo-4/output"
	"os"

	"github.com/fatih/color"
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
		output.PrintError(err)
		return
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}

	color.Green("Запись успешна")
}
