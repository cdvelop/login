package login

// FORM retorna formulario de loging
func FORM() *string {
	h := `
<!-- <div class=" fullBox fullcentrado"> -->
<div id="login" class="panel">
	<h1 class="titulo">Ingreso:</h1>
	<div class="fullcont fullcentrado">
		<div class="loginBox">
			</form name="loging_form">
			<fieldset>
				<legend><label for="usr">Nombre</label></legend>
				<input type="text" value="luis" id="usr">
			</fieldset>

			<fieldset>
				<legend><label for="pwd">Contraseña</label></legend>
				<input type="password" value="123" id="pwd">
			</fieldset>
			</form>
			<div class="cont-btn-login">
				<button name="btnKeyEditlogin" title="Ingresar"></button>
			</div>
		</div>
	</div>
</div>

`
	return &h
}

//JS retorna formulario de loging
func JS() *string {
	//funciones js componente
	j := `
<script>
	const loging = document.querySelector('[name="btnKeyEditlogin"]');
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
