package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CreateUsuario cria um novo usuário
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usuario usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	// PREPARE STATEMENT
	stmt, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	// EXECUTE STATEMENT
	insert, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// GET LAST INSERT ID
	lastID, err := insert.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário inserido com sucesso! ID: " + fmt.Sprint(lastID)))

}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lista de usuários"))
}
