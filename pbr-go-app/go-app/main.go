package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cadastros/", CadastroHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
