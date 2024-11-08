package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaccion struct {
	ID     string  `json:"id"`
	Nombre string  `json:"nombre"`
	Monto  float64 `json:"monto"`
	Tipo   string  `json:"tipo"`
}

var transacciones []Transaccion

func main() {
	router := mux.NewRouter()

	transacciones = append(transacciones, Transaccion{
		ID:     "1",
		Nombre: "Sueldo",
		Monto:  9000,
		Tipo:   "INGRESO",
	})

	router.HandleFunc("/transacciones", GetTransacciones).Methods("GET")
	router.HandleFunc("/transacciones/{id}", GetTransaccion).Methods("GET")
	router.HandleFunc("/transacciones", CreateTransaccion).Methods("POST")

	log.Println("Servidor Iniciado")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetTransacciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transacciones)
}

func GetTransaccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range transacciones {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Transaccion{})
}

func CreateTransaccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaccion Transaccion
	_ = json.NewDecoder(r.Body).Decode(&transaccion)
	transacciones = append(transacciones, transaccion)
	json.NewEncoder(w).Encode(transaccion)
}
