package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"planets-api/pkg/planet"
	"strings"
)

type KeywordsFinder struct {
	Token    json.Token
	Decoder  *json.Decoder
	Keywords []string
	Ch       chan string
	Quit     chan bool
}

func GetPlanetsFromSWAPI() []*planet.Planet {
	resp, err := http.Get("https://swapi.dev/api/planets/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	k := &KeywordsFinder{
		Decoder:  json.NewDecoder(resp.Body),
		Keywords: []string{"name", "climate", "terrain"},
		Ch:       make(chan string),
		Quit:     make(chan bool),
	}
	defer close(k.Ch)
	defer close(k.Quit)

	go k.decode()

	return k.unmarshalData()
}

func (k *KeywordsFinder) decode() {
	for {
		t, err := k.Decoder.Token()
		if err != nil {
			break
		}
		k.Token = t
		k.findValues()
	}
	k.Quit <- true
}

func (k *KeywordsFinder) findValues() {
	for _, word := range k.Keywords {
		if strings.Contains(fmt.Sprintf("%v", k.Token), word) {
			t, err := k.Decoder.Token()
			if err != nil {
				log.Fatal(err)
			}
			k.Ch <- fmt.Sprintf("%v", t)
		}
	}
}

func (k *KeywordsFinder) unmarshalData() []*planet.Planet {
	var (
		result []*planet.Planet
		p      *planet.Planet
		i      int
	)
	for {
		select {
		case v := <-k.Ch:
			switch i % 3 {
			case 0:
				p = &planet.Planet{Name: v}
			case 1:
				p.Climate = v
			case 2:
				p.Terrain = v
				result = append(result, p)
			}
			i++
		case <-k.Quit:
			return result
		}
	}
}
