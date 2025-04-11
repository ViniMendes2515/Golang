package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Lima",
	}

	usuario2 := map[string]map[string]string{
		"nome": {
			"nome":      "Lucas",
			"sobrenome": "Lima",
		},
		"curso": {
			"nome":    "Engenharia",
			"periodo": "2",
		},
	}

	fmt.Println(usuario)
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Aquário",
		"data": "20/01",
	}
	fmt.Println(usuario2["signo"]["nome"])

}
