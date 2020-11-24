package login

// HTMLFORM retorna formulario de loging
func HTMLFORM() *string {
	h := `
<!-- <div class=" fullBox fullcentrado"> -->
<div class=" fullcentrado">
	</form name="loging_form">
	<fieldset>
		<legend><label for="usr">Nombre</label></legend>
		<input type="text" value="luis" id="usr">
	</fieldset>

	<fieldset>
		<legend><label for="pwd">Contraseña</label></legend>
		<input type="password" value="123" id="pwd">
	</fieldset>
	<button name="btnKeyEdit" title="Ingresar"></button>
	</form>
</div>
`
	return &h
}
