package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Gato struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	gatoJson := `{"nome":"Kiff","raca":"Pardo"}`

	var g Gato

	// Passa como parametro o slice de bytes do JSON e o ponteiro da variavel que queremos convertes
	if err := json.Unmarshal([]byte(gatoJson), &g); err != nil {
		log.Fatal(err)
	}

	var m = make(map[string]string)

	if err := json.Unmarshal([]byte(gatoJson), &m); err != nil {
		log.Fatal(err)
	}

	fmt.Println(gatoJson)
	fmt.Println(g)
	fmt.Println(m)

}
