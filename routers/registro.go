package routers

import(
    "encoding/json"
	"net/http"
	"github.com/bblanqui/red_social/bd"
	"github.com/bblanqui/red_social/models"
)
/*Registro es la funcion para crear en la base de datos del usuario*/

func Registro  (w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	 
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "error en los datos resibisod"+err.Error(), 400)
		return
	}

	if len(t.Email)==0{
		http.Error(w, "el email esta vacio"+err.Error(), 400)
		return
	}
	if len(t.Password) < 6{
		http.Error(w, "contraseña corta menor de 6"+err.Error(), 400)
		return
	}

	_, encontrado,_:=bd.ChequeoYaexisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "ya existe usuario"+err.Error(), 400)
		return
	}
	_, encontrado,_:=bd.ChequeoYaexisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "ya existe usuario"+err.Error(), 400)
		return
	}

	_,status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "ocuarrio un error al intentar ingresar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se a logrado insertar correctamente"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
}