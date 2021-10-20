package main

import (
	"log"
	"github.com/bblanqui/red_social/handlers"
	"github.com/bblanqui/red_social/bd"
)


func main() {
if bd.ChequeoConection()==0 {
	log.Fatal("error")
	return
}
handlers.Manejadores()
}
