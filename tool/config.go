/**
 * Created by GoLand.
 * User: adrienli@tencent.com
 * Date: 2019/1/24
 * Time: 3:24 PM
 * Desc: Parse config file
 */
package tool

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	//the map which store config file
	configMap map[string][]byte
	//the map which store param file
	paramsMap map[string][]byte
	//the base config path
	configBasePath string
)

func Parse(confPath string) (configKeys []string, err error)  {
	configBasePath = confPath

	file, err := os.Open(filepath.Join(confPath, "config.json"))

	if err != nil {
		return
	}

	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
	file.Close()

	if err != nil{
		return
	}

	var tempData map[string]map[string]interface{}
	err = json.Unmarshal(content, &tempData)

	if err != nil {
		return
	}
	configMap = make(map[string][]byte, len(tempData))
	configKeys = make([]string, 0)
	for k, v := range tempData {
		configMap[k],_ = json.Marshal(v)
		configKeys = append(configKeys, k)
	}
	return
}

func GetItem(name string, ret interface{}) error {
	if _, ok := configMap[name]; !ok {
		return errors.New("item " + name + " not exist")
	}

	return json.Unmarshal(configMap[name], ret)
}

