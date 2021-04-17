package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getPeriodicTableData() map[string]interface{} {
	jsonFile, err := os.Open("data/pt.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var results map[string]interface{}
	json.Unmarshal([]byte(byteValue), &results)

	periodicTable := make(map[string]interface{}, len(results))
	for k, v := range results {
		periodicTable[strings.ToLower(k)] = v
	}
	return periodicTable
}

func printElementData(inputElement string, periodicTable map[string]interface{}) {
	element := strings.ToLower(inputElement)
	if _, ok := periodicTable[element]; ok {
		symbol := periodicTable[element].(map[string]interface{})["Symbol"]
		fmt.Printf("Element: %s (%s)\n\n", element, symbol)
		fmt.Printf("AtomicMass: %f\n", periodicTable[element].(map[string]interface{})["AtomicMass"])
	} else {
		fmt.Println("Element", element, "does not exist (yet)")
	}
}
