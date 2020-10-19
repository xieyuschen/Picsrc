package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadSetting(configFile string) DbSettings{
	jsonFile,err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal(byteValue,&result)
	dbsettings := result["DbSettings"].(map[string]interface{})

	var set DbSettings
	set.Username = dbsettings["Username"].(string)
	set.Password = dbsettings["Password"].(string)
	set.Hostname = dbsettings["Hostname"].(string)
	set.Dbname = dbsettings["Dbname"].(string)
	return set
}
