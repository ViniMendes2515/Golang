package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "admin:Vmc040104!@/prd?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	println("Conexão com o banco de dados estabelecida com sucesso!")

	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)

}
