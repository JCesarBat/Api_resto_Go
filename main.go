package main

import (
	"Nuevo_go/JCesarBat/Nuevo_go/Handler"
	"Nuevo_go/JCesarBat/Nuevo_go/db"
	"Nuevo_go/JCesarBat/Nuevo_go/model"
	"log"
	"net/http"
)

func main() {
	db.Connection()
	db.DB.AutoMigrate(model.Comida{})

	mux := http.NewServeMux()
	mux.Handle("/Comida/", &Handler.Manejador_Comida{})
	mux.Handle("/Comida", &Handler.Manejador_Comida{})
	log.Fatal(http.ListenAndServe(":3000", mux))

}
