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

type BinList struct {
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

	binInput := Bin{ID: args[1], Private: parsePrivate, CreateAt: createAtParse, Name: args[4]}

	fmt.Println(binInput)

}
