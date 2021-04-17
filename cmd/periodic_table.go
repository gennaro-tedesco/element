package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed pt.json
var b []byte

func getPeriodicTableData() map[string]interface{} {
	var results map[string]interface{}
	json.Unmarshal([]byte(b), &results)

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
		phase := periodicTable[element].(map[string]interface{})["Phase"]
		fmt.Printf("Element: %s (%s) - %s \n\n", element, symbol, phase)
		fmt.Printf("AtomicMass: %v\n", periodicTable[element].(map[string]interface{})["AtomicMass"])
		fmt.Printf("AtomicNumber: %v\n", periodicTable[element].(map[string]interface{})["AtomicNumber"])
		fmt.Printf("Density: %v\n", periodicTable[element].(map[string]interface{})["Density"])
		fmt.Printf("Protons: %v\n", periodicTable[element].(map[string]interface{})["NumberofProtons"])
		fmt.Printf("Electrons: %v\n", periodicTable[element].(map[string]interface{})["NumberofElectrons"])
		fmt.Printf("Neutrons: %v\n", periodicTable[element].(map[string]interface{})["NumberofNeutrons"])
		fmt.Printf("AtomicRadius: %v\n", periodicTable[element].(map[string]interface{})["AtomicRadius"])
		fmt.Printf("SpecificHeat: %v\n", periodicTable[element].(map[string]interface{})["SpecificHeat"])
		fmt.Printf("Electronegativity: %v\n", periodicTable[element].(map[string]interface{})["Electronegativity"])
	} else {
		fmt.Println("Element", element, "does not exist (yet)")
	}
}

func getRandomKey(periodicTable map[string]interface{}) string {
	for k := range periodicTable {
		return k
	}
	return "hydrogen"
}
