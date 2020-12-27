package login

import (
	"fmt"
	"net/http"
)

//INGRESO acceso al sistema
func INGRESO(w http.ResponseWriter, r *http.Request) {
	resp := "metodo de conexion no permitida" //respuesta servidor por defecto
	if r.Method == http.MethodPost {

		usr, _ := r.URL.Query()["usr"] //user
		pwd, _ := r.URL.Query()["pwd"] //password

		// buscar usuario en db
		if p, x := usuarios[usr[0]]; x && p["pwd"] == pwd[0] {

			resp = fmt.Sprintf(">>> %v acceso autorizado", usr)
		} else {
			resp = fmt.Sprintf("!!! %v acceso denegado", usr)
		}

	}
	w.Write([]byte(resp))
}
