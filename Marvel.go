// Una _goroutine_ es un hilo o thread de ejecución ligero.
// API: https://developer.marvel.com

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	copyright string `json:"copyright"`
	data      Data   `json:"data"`
	code      string `json:"code"`
	message   string `json:"message"`
}

type Data struct {
	total   int       `json:"total"`
	results []Results `json:"results"`
	count   int       `json:"count"`
}

type Results struct {
	name        string    `json:"name"`
	description string    `json:"description"`
	thumbnail   Thumbnail `json:"thumbnail"`
}

type Thumbnail struct {
	path      string `json:"path"`
	extension string `json:"extension"`
}

func PeticionAPI(limite int, inicio int) {

	limit := strconv.Itoa(limite)
	offset := strconv.Itoa(inicio)

	fmt.Println("El limite es: ", limit)
	fmt.Println("El inicio es: ", offset)

	response, err := http.Get("https://gateway.marvel.com:443/v1/public/characters?limit=" + limit + "&offset=" + offset + "&apikey=06ffa280d1bafc06d930b43d6d8dd14b&hash=afda8720864a69268e1e8bedd7a23b60")
	if err != nil {
		fmt.Println("Primer IF")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Segundo IF")
		log.Fatal(err)
	}

	fmt.Println("Salio IF")

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("code: " + responseObject.code)
	fmt.Println("message: " + responseObject.message)

	for i := 1; i < responseObject.data.count; i++ {
		fmt.Println("Entro en FOR", i)
		fmt.Println(responseObject.data.results[i].name, "\n")
		fmt.Println(responseObject.data.results[i].description, "\n")
		fmt.Println(responseObject.data.results[i].thumbnail.path, ".", responseObject.data.results[i].thumbnail.extension, "\n")
	}
}

func main() {

	var limit int
	var offset int

	for i := 0; i < 10; i++ {
		offset = rand.Intn(1541)
		limit = rand.Intn(20)
		go PeticionAPI(limit, offset)
	}

	var input string
	fmt.Scanln(&input)
}
