/**
 * Created by GoLand.
 * User: adrienli@tencent.com
 * Date: 2019/1/24
 * Time: 3:06 PM
 * Desc:
 */
package main

import (
	"Weather/api"
	"Weather/tool"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		key        string
		city       string
		extensions string
		output     string
	)

	defer func() {
		i := recover()
		if i != nil {
			return
		}
	}()

	city = os.Args[1]
	key = tool.API_KEY
	extensions = tool.EXTENSIONS
	output = tool.OUTPUT
	city = strings.Trim(city, "\n")
	weatherInfo, err := api.GetWeather(key, city, extensions, output)

	if err != nil {
		panic(err)
	}

	for _, v := range weatherInfo.Lives {
		fmt.Println("省份:", v.Province)
		fmt.Println("城市:", v.City)
		fmt.Println("天气:", v.Weather)
		fmt.Println("气温:", v.Temperature)
		fmt.Println("风向:", v.WindDirection)
		fmt.Println("风力:", v.WindPower)
		fmt.Println("空气湿度:", v.Humidity)
		fmt.Println("发布时间:", v.ReportTime)
	}
}
