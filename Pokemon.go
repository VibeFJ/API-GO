// Una _goroutine_ es un hilo o thread de ejecución ligero.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func PeticionAPI(i int) {

	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/original-alola")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Pokemon[i].Species.Name, "\n")
}

func main() {

	var Random int

	for i := 0; i < 10; i++ {
		Random = rand.Intn(151)
		go PeticionAPI(Random)
	}

	var input string
	fmt.Scanln(&input)
}
