package main

import (
	"bytes"
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
	gato := Gato{
		Nome:  "Miau",
		Raca:  "Siamês",
		Idade: 3,
	}

	jsonGato, err := json.Marshal(gato)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonGato)
	fmt.Println(bytes.NewBuffer(jsonGato))

	gato2 := map[string]string{
		"nome": "Kiff",
		"raca": "Pardo",
	}

	jsonGato2, err := json.Marshal(gato2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonGato2)
	fmt.Println(bytes.NewBuffer(jsonGato2))

}
