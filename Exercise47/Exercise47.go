package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
   Create a program that pulls in data from the Open Notify API that
   provides information on how many people and who is currently in space,
   and which craft they are on

   Constraint:

   Read the data directly from the API and parse the results each time the
   program is run. Don't download the data as text and read it in.
*/

type SpaceData struct {
	Message string      `json:"message"`
	Number  int         `json:"number"`
	People  []Astronaut `json:"people"`
}

type Astronaut struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

func GrabData() SpaceData {
	var sd SpaceData
	resp, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		log.Fatalf("Error reading from API: %s\n", err.Error())
		return sd
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading API response body: %s\n", err.Error())
	}
	if err := json.Unmarshal(data, &sd); err != nil {
		log.Fatalf("JSON conversion Error: %s\n", err.Error())
	}
	return sd
}

func displayAstronauts(sd SpaceData) {
	fmt.Printf("Data Processing: %s\n", sd.Message)
	fmt.Printf("There are currently %d people in space\n", sd.Number)
	fmt.Printf("%-25s | %-10s\n", "Name", "Craft")
	fmt.Println("------------------------- | ----------")
	for _, astro := range sd.People {
		fmt.Printf("%-25s | %-10s\n", astro.Name, astro.Craft)
	}
}

func main() {
	s := GrabData()
	displayAstronauts(s)
}
