/**
 * Created by GoLand.
 * User: adrienli@tencent.com
 * Date: 2019/1/24
 * Time: 3:11 PM
 * Desc:
 */
package api

import (
	"Weather/data"
	"Weather/tool"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func GetWeather(key, city, extensions, output string) (weatherInfo data.ApiResponse, err error){
	var (
		url     string
		curlUrl string
	)

	url = tool.WEATHER_URL
	curlUrl = fmt.Sprint(url,
		"?key=", key,
		"&city=", city,
		"&extensions=", extensions,
		"&output=", output,
	)
	fmt.Println(curlUrl)
	response, err := http.Get(curlUrl)
	fmt.Println("response", response)
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		msg := "get weather info from api failed"
		errors.New(msg)
		panic(err)
	}

	err = json.Unmarshal(responseBody, &weatherInfo)
	return
}
