package main

import (
	"flag"
	"fmt"
	"go-demo-5/geo"
)

func main() {
	city := flag.String("city", "", "Город пользователя")

	flag.Parse()

	fmt.Println("city--->", *city)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("geoData-->",geoData)

}
