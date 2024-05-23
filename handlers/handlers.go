package handlers

import (
	"context"
	"fmt"

	"github.com/MarcoABCH/GoAPITwitter/jwt"
	"github.com/MarcoABCH/GoAPITwitter/models"
	"github.com/MarcoABCH/GoAPITwitter/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi{
	fmt.Println("Voy a procesar "+ ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	//Default hay un error
	r.Status=400

	isOK, statusCode, msg, _ := validoAuthorization(ctx, request)//el 4to parametro es claim pero la quitamos con _ 

	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r	
	}


	switch ctx.Value(models.Key("method")).(string) {//Evaluamos el tipo de metodo(Verbo de la api)
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {//Evaluamos el endpoint(accion) de la api
		case "registro":
			return routers.Registro(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {//Evaluamos el endpoint(accion) de la api
		
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {//Evaluamos el endpoint(accion) de la api
		
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {//Evaluamos el endpoint(accion) de la api
		
		}
	}

	r.Message="Method Invalid"
	return r
}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim){
	path:= ctx.Value(models.Key("path")).(string)
	if path =="registro" || path =="login" || path =="obtenerAvatar" || path =="obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token:= request.Headers["Authorization"]
	if len(token) == 0 {
		return true, 401, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOK {
		if err!= nil {
			fmt.Println("Error en el token "+err.Error())
			return false, 401, err.Error(), models.Claim{}
		}else {
			fmt.Println("Error en el token "+msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}