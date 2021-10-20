package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bblanqui/red_social/middlew"
	"github.com/bblanqui/red_social/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo el handler y pongo a escuchar el srvidor*/
func Manejadores() {

	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {

		PORT = os.Getenv("8000")
	}
	handler := cors.AllowAll().Handler(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func holamundo (w http.ResponseWriter, r *http.Request){
     fmt.Println("hola munfo")
}