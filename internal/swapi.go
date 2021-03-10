package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RespBody struct {
	Results []Results
}

type Results struct {
	Name    string
	Climate string
	Terrain string
	Films   []string
}

func main() {
	resp, err := http.Get("https://swapi.dev/api/planets")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var body RespBody
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Fatal(err)
	}
	for _, p := range body.Results {
		fmt.Println(p)
	}
}

func (r Results) String() string {
	return fmt.Sprintf(
		"\nPLANET\nName: %s\nClimate: %s\nTerrain: %s\nFilms: %v",
		r.Name, r.Climate, r.Terrain, len(r.Films),
	)
}
