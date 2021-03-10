package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"planets-api/pkg/planet"
	"testing"

	"github.com/stretchr/testify/assert"
)

const url = "http://localhost:8080"

func Test_GetAllPlanets(t *testing.T) {
	var test = planet.Planet{Name: "Yavin 4", Climate: "Tropical", Terrain: "Forests"}
	// Create
	add, err := receiveResp(doRequest("POST", url+"/planet", test))
	assert.NoError(t, err)
	assert.Equal(t, test.Name, add.Name)
	assert.Equal(t, test.Climate, add.Climate)
	assert.Equal(t, test.Terrain, add.Terrain)

	// GetAll
	out := doRequest("GET", url+"/planet", test)
	var p []struct {
		ID      string
		Name    string
		Terrain string
		Climate string
	}
	if err := json.Unmarshal(out.Bytes(), &p); err != nil {
		log.Fatalln(err, out.String())
	}
}

func Test_getPlanetWithName(t *testing.T) {
	var test = planet.Planet{Name: "Naboo", Climate: "Tropical", Terrain: "Forests"}

	// Create
	add, err := receiveResp(doRequest("POST", url+"/planet", test))
	assert.NoError(t, err)
	assert.Equal(t, test.Name, add.Name)
	assert.Equal(t, test.Climate, add.Climate)
	assert.Equal(t, test.Terrain, add.Terrain)

	// Get
	get, err := receiveResp(doRequest(
		"GET", fmt.Sprintf("%s/planet/%s", url, test.Name), test))
	assert.NoError(t, err)
	assert.Equal(t, add, get)
}

func Test_getPlanetWithID(t *testing.T) {
	var test = planet.Planet{Name: "Mon Calamari", Climate: "Fresh", Terrain: "Water"}

	// Create
	add, err := receiveResp(doRequest("POST", url+"/planet", test))
	assert.NoError(t, err)
	assert.Equal(t, test.Name, add.Name)
	assert.Equal(t, test.Climate, add.Climate)
	assert.Equal(t, test.Terrain, add.Terrain)

	// Get
	get, err := receiveResp(doRequest(
		"GET", fmt.Sprintf("%s/planet/%v", url, add.ID.Hex()), test))
	assert.NoError(t, err)
	assert.Equal(t, add, get)
}

func Test_deletePlanet(t *testing.T) {
	var test = planet.Planet{Name: "test", Climate: "Cold", Terrain: "Ice and snow"}

	// Create
	add, err := receiveResp(doRequest("POST", url+"/planet", test))
	assert.NoError(t, err)
	assert.Equal(t, test.Name, add.Name)
	assert.Equal(t, test.Climate, add.Climate)
	assert.Equal(t, test.Terrain, add.Terrain)

	// Delete
	route := fmt.Sprintf("/planet/%v", add.ID.Hex())
	del := doRequest("DELETE", url+route, test)
	assert.Equal(t, add.ID.Hex(), del.String())
}

func Test_addPlanet(t *testing.T) {
	var test = planet.Planet{Name: "Mustafar", Climate: "Hot", Terrain: "Volcanoes and lava"}

	resp, err := receiveResp(doRequest("POST", url+"/planet", test))
	assert.NoError(t, err)
	assert.Equal(t, test.Name, resp.Name)
	assert.Equal(t, test.Climate, resp.Climate)
	assert.Equal(t, test.Terrain, resp.Terrain)
}

func doRequest(method, route string, p planet.Planet) *bytes.Buffer {
	var cmd *exec.Cmd
	switch method {
	case "GET":
		cmd = exec.Command("curl", "-X", "GET", route)
	case "POST":
		cmd = exec.Command(
			"curl",
			"-X", "POST",
			"-d", fmt.Sprintf("name=%v", p.Name),
			"-d", fmt.Sprintf("climate=%v", p.Climate),
			"-d", fmt.Sprintf("terrain=%v", p.Terrain),
			route, // http address
		)
	case "DELETE":
		cmd = exec.Command("curl", "-X", "DELETE", route)
	}
	out := &bytes.Buffer{}
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	return out
}

func receiveResp(out *bytes.Buffer) (planet.Planet, error) {
	var p planet.Planet
	if err := json.Unmarshal(out.Bytes(), &p); err != nil {
		return p, err
	}
	return p, nil
}
