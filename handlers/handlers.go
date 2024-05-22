package handlers

import (
	"context"
	"fmt"

	"github.com/MarcoABCH/GoAPITwitter/models"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RestApi{
	fmt.Println("Voy a procesar "+ ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RestApi

	//Default hay un error
	r.Status=400

	switch ctx.Value(models.Key("method")).(string) {//Evaluamos el tipo de metodo(Verbo de la api)
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {//Evaluamos el endpoint(accion) de la api
		
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