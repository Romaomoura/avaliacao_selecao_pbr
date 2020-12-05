package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	//"math/big"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/lxn/go-pgsql"

)

// Cadastro
type Cadastro struct {
	ID   int `json:"id"`
	NomeCompleto string `json:"nomeCompleto"`
	Telefone string `json:"telefone"`
	Email string `json:"email"`
}

// UsuarioHandler analisa o request e delega para função adequada
func CadastroHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/cadastros/")
	id, _ := strconv.Atoi(sid)
	fmt.Println(id) //verificar se o id é o mesmo repassado na URL

	switch {
	case r.Method == "GET" && id > 0:
		//fmt.Println(id)
		buscarPorID(w, r, id)
	case r.Method == "GET":
		buscarTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Não deu... :(")
	}
}
//Configuração básica do banco de dados, dados expostos (forma incorreta)
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "Romao1988!"
	DB_NAME     = "api_pbr"
)


/* func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
} */

func buscarPorID(w http.ResponseWriter, r *http.Request, id int) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	idLong := uint64(id)

	var cad Cadastro
	db.QueryRow("SELECT id, nome_completo, telefone, email FROM cadastro WHERE id=?", idLong).Scan(&cad.ID, &cad.NomeCompleto, &cad.Telefone, &cad.Email)
	fmt.Println(idLong) //verifica qual id está sendo passado
	fmt.Println(cad.ID, cad.NomeCompleto) //verificar se a consulta está retornando algum dado

	json, _ := json.Marshal(cad)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func buscarTodos(w http.ResponseWriter, r *http.Request) {
	
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, email, nome_completo, telefone from cadastro")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cadastros []Cadastro
	for rows.Next() {
		var cadastro Cadastro
		rows.Scan(&cadastro.ID, &cadastro.Email, &cadastro.NomeCompleto, &cadastro.Telefone)
		cadastros = append(cadastros, cadastro)
	}

	json, _ := json.Marshal(cadastros)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
