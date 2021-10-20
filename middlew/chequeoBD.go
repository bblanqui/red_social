package middlew

import (
	"net/http"

	"github.com/bblanqui/red_social/bd"
)

func ChequeoBD (next http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
       if bd.ChequeoConection()== 0 {
		   http.Error(w, "conexion perdida", 500)
		   return
	   }
	   next.ServeHTTP(w,r)
	}
}