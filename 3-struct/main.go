package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Bin struct {
	ID       string
	Private  bool
	CreateAt time.Time
	Name     string
}

func NewBin(id string, private bool, createAt time.Time, name string) Bin {
	return Bin{
		ID:       id,
		Private:  private,
		CreateAt: createAt,
		Name:     name,
	}
}

type BinList struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBinList(id string, private bool, createdAt time.Time, name string) BinList {
	return BinList{
		ID:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

func main() {
	args := os.Args

	if len(args) != 5 {
		fmt.Println("Ошибка")
		return
	}

	parsePrivate, err := strconv.ParseBool(args[2])
	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	createAtStr := args[3]
	createAtParse, err := time.Parse("2006-01-02T15:04:05", createAtStr)
	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	binInput := NewBin(args[1], parsePrivate, createAtParse, args[4])
	fmt.Println(binInput)
}
