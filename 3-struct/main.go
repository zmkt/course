package main

import (
	"3-struct/api"
	"3-struct/bins"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args

	if len(args[1]) != 5 {
		fmt.Println("Ошибка1")
		return
	}

	parsePrivate, err := strconv.ParseBool(args[2])
	if err != nil {
		fmt.Println("Ошибка2")
		return
	}

	createAtStr := args[3]
	createAtParse, err := time.Parse("2006-01-02T15:04:05", createAtStr)
	if err != nil {
		fmt.Println("Ошибка3")
		return
	}

	apiCfg, err := api.NewConfig()
	if err != nil {
		fmt.Println("Ошибка конфига:", err)
		return
	}

	binInput := bins.NewBin(args[1], parsePrivate, createAtParse, args[4])

	fmt.Println("API key:", apiCfg.Key)
	fmt.Println(binInput)
}
