/**
 * Created by GoLand.
 * User: adrienli@tencent.com
 * Date: 2019/1/24
 * Time: 3:34 PM
 * Desc:
 */
package tool

import (
	"fmt"
	"testing"
)

type apiRet struct {
	Api keyRet `json:"api"`
}

type keyRet struct {
	Key string `json:"key"`
}

func TestParse(t *testing.T) {
	t.Run("test Parse", func(t *testing.T) {
		Parse("../config/")

		var config keyRet
		ret := GetItem("api", &config)
		fmt.Println("api", config, ret)
	})
}
