package login

// HTMLJS retorna formulario de loging
func HTMLJS() *string {
	//funciones js componente
	j := `
<script>
	const loging = modulo.querySelector('[name="btnKeyEdit"]');
	loging.onclick = loginServer;

	function loginServer(evt) {
		evt.preventDefault();
		usr = document.querySelector('#usr').value;
		pwd = document.querySelector('#pwd').value;

		conecServer();

	}

</script>`
	return &j
}
