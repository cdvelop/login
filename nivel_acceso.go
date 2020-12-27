package login

import "strings"

//GETMODNVLA obtener array modulos segun key 3 caracteres
func GETMODNVLA(op string, NVLA *map[string][]string) (m []string) {
	for kn, ar := range *NVLA {
		if strings.Contains(strings.ToLower(kn[0:3]), strings.ToLower(op)) {
			m = ar
			return
		}
	}
	return
}

//retona seleccion des,adm,psi|traduccionk Desarrolador| []string key abreviado modulos
// filtro modulos por nivel

//NATOKEY transforma los key de nivel de acceso a minusculas las primeras 3 letras
func NATOKEY(NVLA *map[string][]string) (k []string) {
	//abreviatura 3 primeras letras  a minusculas
	for op := range *NVLA {
		k = append(k, strings.ToLower(op[0:3]))
	}
	return
}
