package login

//cv propiedades tipo clave: valor
type cv map[string]string

var usuarios = make(map[string]map[string]string)

func init() {

	usuarios = map[string]map[string]string{
		"juan":  cv{"pwd": "123", "nvl": "1"},
		"maria": cv{"pwd": "123", "nvl": "1"},
		"luis":  cv{"pwd": "123", "nvl": "2"},
	}

}

//me modelos end point (rutas) staticas del sistema
func me() *map[string]map[string]string {
	//ejemplo de rutas
	e := map[string]map[string]string{
		"hora":     cv{"nvl": "1"},
		"paciente": cv{"nvl": "2"},
	}
	return &e
}
