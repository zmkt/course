package weather

import (
	"errors"
	"fmt"
	"go-demo-5/geo"
	"io"
	"net/http"
	"net/url"
)

var Error200 = errors.New("NOT200")

func GetWeather(geo geo.GeoData, format int) (string, error) {

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", Error200
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", Error200
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", Error200
	}

	return string(body), nil

}
