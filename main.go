package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.83
	userKg := 95.0
	userHeight = 1.83
	var IMT = userKg / math.Pow(userHeight, IMTPower)

	fmt.Println(IMT)
}