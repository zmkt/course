package main

import (
	"3-struct/bins"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

	binInput := bins.NewBin(args[1], parsePrivate, createAtParse, args[4])
	fmt.Println(binInput)
}
