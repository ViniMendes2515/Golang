package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com o MySQL
)

// Connect conecta ao banco de dados MySQL e retorna a conexão
func Connect() (*sql.DB, error) {
	stringConnection := "admin:Vmc040104!@/prd?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
