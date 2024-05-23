package main

import (
	"context"
	"os"
	"strings"

	"github.com/MarcoABCH/GoAPITwitter/awsgo" //Importamos nuestro archivo de comunicacion con AWS
	"github.com/MarcoABCH/GoAPITwitter/bd"
	"github.com/MarcoABCH/GoAPITwitter/handlers"
	"github.com/MarcoABCH/GoAPITwitter/models"
	"github.com/MarcoABCH/GoAPITwitter/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)


func main(){
	lambda.Start(EjecutoLambda)
}
//Para llamar el API Gateway de AWS, esta funcion desde aws nos va retornar una info de tipo puntero
func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error){
	var res *events.APIGatewayProxyResponse

	awsgo.InicializoAWS()

	if !ValidoParametros(){
		res = &events.APIGatewayProxyResponse {
			StatusCode: 400,
			Body: "Error en las variables de entorno, deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}	
		return res, nil	
	}	

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
			
	if err != nil {
		res = &events.APIGatewayProxyResponse {
			StatusCode: 400,
			Body: "Error en en la lectura de Secret "+ err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}	
		return res, nil	
	}

	path := strings.Replace(request.PathParameters["goapitwitter"], os.Getenv("UrlPrefix"),"", -1) 
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.DataBase)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//Checar conexion a la bd
	err = bd.ConectarBD(awsgo.Ctx)
	if err!= nil {
		res = &events.APIGatewayProxyResponse {
			StatusCode: 500,
			Body: "Error al conectar a la BD "+ err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}	
		return res, nil	
	}

	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	//Si despues de la peticion no viene bien armada la respuesta la generamos una respuesta personalizamos
	if respAPI.CustomResp == nil{
		res = &events.APIGatewayProxyResponse {
			StatusCode: respAPI.Status,
			Body: respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}	

		return res, nil
	}else{
		return respAPI.CustomResp, nil
	}
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")

	if !traeParametro {
		return traeParametro
	}

	//Este es para quitar el prefijo que trae la ruta y dejar bien limpio el llamado
	_, traeParametro = os.LookupEnv("UrlPrefix")

	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}