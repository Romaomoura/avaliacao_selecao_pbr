package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

/* func main() {
	http.HandleFunc("/cadastros/", CadastroHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
} */

// Cadastro
type Cadastro struct {
	ID   int    `json:"id"`
	NomeCompleto string `json:"nomeCompleto"`
	Telefone string `json:"telefone"`
	Email string `json:"email"`
}

// UsuarioHandler analisa o request e delega para função adequada
func CadastroHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/cadastros/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		buscarPorID(w, r, id)
	case r.Method == "GET":
		buscarTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Não deu... :(")
	}
}
//Configuração básica do banco de dados
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
	//db, err := sql.Open("postgres", "postgres:Romao1988!@/api_pbr")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var c Cadastro
	db.QueryRow("select id, email, nome_completo, telefone from cadastro where id = ?", id).Scan(&c.ID, &c.Email, &c.NomeCompleto, &c.Telefone)

	json, _ := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func buscarTodos(w http.ResponseWriter, r *http.Request) {
	
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	//db, err := sql.Open("postgres", "postgres:Romao1988!@/api_pbr")
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
