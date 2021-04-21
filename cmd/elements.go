package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

//go:embed pt.json
var ptFile []byte

func getPeriodicTableData() map[string]interface{} {
	var results map[string]interface{}
	json.Unmarshal([]byte(ptFile), &results)

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
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Property", "Value"})
		t.AppendRows([]table.Row{
			{"Protons", fmt.Sprint(periodicTable[element].(map[string]interface{})["NumberofProtons"])},
			{"Electrons", fmt.Sprint(periodicTable[element].(map[string]interface{})["NumberofElectrons"])},
			{"Neutrons", fmt.Sprint(periodicTable[element].(map[string]interface{})["NumberofNeutrons"])},
			{"Period", fmt.Sprint(periodicTable[element].(map[string]interface{})["Period"])},
			{"Group", fmt.Sprint(periodicTable[element].(map[string]interface{})["Group"])},
			{"Electronic configuration", fmt.Sprint(periodicTable[element].(map[string]interface{})["ElectronicConfiguration"])},
			{"Atomic mass", fmt.Sprint(periodicTable[element].(map[string]interface{})["AtomicMass"])},
			{"Atomic number", fmt.Sprint(periodicTable[element].(map[string]interface{})["AtomicNumber"])},
			{"Atomic radius", fmt.Sprint(periodicTable[element].(map[string]interface{})["AtomicRadius"])},
			{"Density (g/cm^3)", fmt.Sprint(periodicTable[element].(map[string]interface{})["Density"])},
			{"Specific heat (J/K)", fmt.Sprint(periodicTable[element].(map[string]interface{})["SpecificHeat"])},
			{"Electronegativity", fmt.Sprint(periodicTable[element].(map[string]interface{})["Electronegativity"])},
			{"Melting point (K)", fmt.Sprint(periodicTable[element].(map[string]interface{})["MeltingPoint"])},
			{"Boiling point (K)", fmt.Sprint(periodicTable[element].(map[string]interface{})["BoilingPoint"])},
			{"First ionization (eV)", fmt.Sprint(periodicTable[element].(map[string]interface{})["FirstIonization"])},
		})
		t.AppendSeparator()
		t.SetStyle(table.StyleLight)
		t.Render()

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
