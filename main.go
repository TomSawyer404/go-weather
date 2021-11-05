package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const APPID = `4845f22236e074cdac59ae174aa580a3`
const RED = "\x1b[31m"
const RESET = "\x1b[0m"

type (
	W struct {
		Main    Main
		Weather Weather
	}

	Main struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
		Pressure   int32   `json:"pressure"`
		Humidity   int32   `json:"humidity"`
	}

	Weather []struct {
		Id          uint32 `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
)

func main() {
	argv := os.Args
	if len(argv) != 2 {
		log.Fatalf("Usage: %s <CITY_NAME>\n", argv[0])
	}

	u := &url.URL{
		Scheme:   `http`,
		Host:     `api.openweathermap.org`,
		Path:     `data/2.5/weather`,
		RawQuery: fmt.Sprintf("q=%s&appid=%s", argv[1], APPID),
	}
	resp, err := http.DefaultClient.Get(u.String())
	if err != nil {
		log.Fatalln(`http.DefaultClient.Get() ->`, err)
	}
	defer resp.Body.Close()

	print_respon(resp)
}

func print_respon(resp *http.Response) {
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(`ioutil.ReadAll() ->`, err)
	}

	var my_weather W
	if err := json.Unmarshal(resp_body, &my_weather); err != nil {
		log.Fatalln(`json.Unmarshal() ->`, err)
	}

	//fmt.Println(string(resp_body))
	//fmt.Println(my_weather)

	fmt.Printf("Weather: %s%s%s\n", RED, my_weather.Weather[0].Main, RESET)
	fmt.Printf("Description: %s%s%s\n", RED, my_weather.Weather[0].Description, RESET)
	fmt.Printf("Current Tempreature: %s%.2f Celsius%s\n", RED, my_weather.Main.Temp-273, RESET)
	fmt.Printf("Feels like: %s%.2f Celsius%s\n", RED, my_weather.Main.Feels_like-273, RESET)
	fmt.Printf("Humidity: %s%d%s\n", RED, my_weather.Main.Humidity, RESET)
	fmt.Printf("Pressure: %s%d Pascal%s\n", RED, my_weather.Main.Pressure, RESET)
}
