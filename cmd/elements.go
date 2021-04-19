package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
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
		elementType := periodicTable[element].(map[string]interface{})["Type"]
		fmt.Printf("Element: %s (%s) - %s, %s \n\n", element, symbol, phase, strings.ToLower(elementType.(string)))
		fmt.Printf("Protons: %v\n", periodicTable[element].(map[string]interface{})["NumberofProtons"])
		fmt.Printf("Electrons: %v\n", periodicTable[element].(map[string]interface{})["NumberofElectrons"])
		fmt.Printf("Neutrons: %v\n\n", periodicTable[element].(map[string]interface{})["NumberofNeutrons"])
		fmt.Printf("Period: %v\n", periodicTable[element].(map[string]interface{})["Period"])
		fmt.Printf("Group: %v\n", periodicTable[element].(map[string]interface{})["Group"])
		fmt.Printf("ElectronicConfiguration: %v\n\n", periodicTable[element].(map[string]interface{})["ElectronicConfiguration"])
		fmt.Printf("AtomicMass: %v\n", periodicTable[element].(map[string]interface{})["AtomicMass"])
		fmt.Printf("AtomicNumber: %v\n", periodicTable[element].(map[string]interface{})["AtomicNumber"])
		fmt.Printf("AtomicRadius: %v\n", periodicTable[element].(map[string]interface{})["AtomicRadius"])
		fmt.Printf("Density: %v\n", periodicTable[element].(map[string]interface{})["Density"])
		fmt.Printf("SpecificHeat: %v\n", periodicTable[element].(map[string]interface{})["SpecificHeat"])
		fmt.Printf("Electronegativity: %v\n", periodicTable[element].(map[string]interface{})["Electronegativity"])
		fmt.Printf("MeltingPoint: %v\n", periodicTable[element].(map[string]interface{})["MeltingPoint"])
		fmt.Printf("BoilingPoint: %v\n", periodicTable[element].(map[string]interface{})["BoilingPoint"])
		fmt.Printf("FirstIonization: %v\n", periodicTable[element].(map[string]interface{})["FirstIonization"])
	} else if element == "" {
		os.Exit(1)
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

func printPeriodicTable() {
	asciiTable := `
   1A   2A                                         3A  4A  5A  6A  7A  8A
  -----                                                               -----
1 | H |                                                               |He |
  |---+----                                       --------------------+---|
2 |Li |Be |                                       | B | C | N | O | F |Ne |
  |---+---|                                       |---+---+---+---+---+---|
3 |Na |Mg |3B  4B  5B  6B  7B |    8B     |1B  2B |Al |Si | P | S |Cl |Ar |
  |---+---+---------------------------------------+---+---+---+---+---+---|
4 | K |Ca |Sc |Ti | V |Cr |Mn |Fe |Co |Ni |Cu |Zn |Ga |Ge |As |Se |Br |Kr |
  |---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
5 |Rb |Sr | Y |Zr |Nb |Mo |Tc |Ru |Rh |Pd |Ag |Cd |In |Sn |Sb |Te | I |Xe |
  |---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
6 |Cs |Ba |LAN|Hf |Ta | W |Re |Os |Ir |Pt |Au |Hg |Tl |Pb |Bi |Po |At |Rn |
  |---+---+---+------------------------------------------------------------
7 |Fr |Ra |ACT|
  -------------------------------------------------------------------------
   Lanthanide |La |Ce |Pr |Nd |Pm |Sm |Eu |Gd |Tb |Dy |Ho |Er |Tm |Yb |Lu |
              |---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
   Actinide   |Ac |Th |Pa | U |Np |Pu |Am |Cm |Bk |Cf |Es |Fm |Md |No |Lw |
              -------------------------------------------------------------
`
	fmt.Println(asciiTable)
}
