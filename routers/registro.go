package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MarcoABCH/GoAPITwitter/bd"
	"github.com/MarcoABCH/GoAPITwitter/models"
)

func Registro(ctx context.Context) models.RestApi {
	var t models.Usuario
	var r models.RestApi
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err!=nil{
		r.Message=err.Error()
		fmt.Println(r.Message)
		return r
	}
	
	if len(t.Email) == 0 {
		r.Message="Debe de especificar el email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message="Debe de especificar el password de al menos 6 carcateres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message="Ya existe un usuario registrado con ese email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err:= bd.InsertoRegistro(t)
	if err!=nil{
		r.Message="Ocurrio un error al registrar el usuario "+err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status{
		r.Message="No se ha logrado insertar el usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status=200
	r.Message="Registro Ok"		
	return r

}