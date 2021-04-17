package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func getPeriodicTableData() map[string]interface{} {
	jsonFile, err := os.Open("data/pt.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var periodicTable map[string]interface{}
	json.Unmarshal([]byte(byteValue), &periodicTable)
	return periodicTable
}
