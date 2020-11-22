package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadSetting(configFile string) JsonSettings {
	jsonFile,err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf JsonSettings
	json.Unmarshal(byteValue,&conf)
	return conf
}
