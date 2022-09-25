package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"strings"
)

type Dependency struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

func main() {
	jsonFossa, _ := ioutil.ReadFile("sca-output.json")

	//xu ly sourceUnits
	sourceUnitsDependencies := gjson.Get(string(jsonFossa), "sourceUnits.#.Build.Dependencies")
	sourceUnitsImports := gjson.Get(string(jsonFossa), "sourceUnits.#.Build.Imports")
	projectsGraphDeps := gjson.Get(string(jsonFossa), "projects.#.graph.deps")

	var dependencies []Dependency
	for _, sourceUnitsDependency := range sourceUnitsDependencies.Array() {
		for _, item := range sourceUnitsDependency.Array() {
			locator := item.Get("locator").String() //string
			indexPlus := strings.Index(locator, "+")
			indexDola := strings.Index(locator, "$")
			libType := locator[0:indexPlus]
			libName := locator[indexPlus+1 : indexDola]
			libVersion := locator[indexDola+1 : len(locator)]

			newDependency := Dependency{Name: libName, Type: libType, Version: libVersion}
			dependencies = append(dependencies, newDependency)
		}
	}

	for _, sourceUnitsImport := range sourceUnitsImports.Array() {
		for _, item := range sourceUnitsImport.Array() {
			itemString := item.String()
			indexPlus := strings.Index(itemString, "+")
			indexDola := strings.Index(itemString, "$")
			libType := itemString[0:indexPlus]
			libName := itemString[indexPlus+1 : indexDola]
			libVersion := itemString[indexDola+1 : len(itemString)]

			newDependency := Dependency{Name: libName, Type: libType, Version: libVersion}
			dependencies = append(dependencies, newDependency)
		}
	}

	for _, projectsGraphDep := range projectsGraphDeps.Array() {
		for _, item := range projectsGraphDep.Array() {
			libType := item.Get("type").String()
			libName := item.Get("name").String()
			libVersion := item.Get("version.value").String()

			newDependency := Dependency{Name: libName, Type: libType, Version: libVersion}
			fmt.Println(newDependency)
			dependencies = append(dependencies, newDependency)
		}
	}

	file, _ := json.MarshalIndent(dependencies, "", " ")

	_ = ioutil.WriteFile("sca-parser.json", file, 0644)

	b, err := json.Marshal(dependencies)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
