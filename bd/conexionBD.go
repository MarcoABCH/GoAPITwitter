package bd

import (
	"context"
	"fmt"

	"github.com/MarcoABCH/GoAPITwitter/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var DatabaseName string

func ConectarBD(ctx context.Context) error {
	user :=ctx.Value(models.Key("user")).(string)
	paswd :=ctx.Value(models.Key("password")).(string)
	host :=ctx.Value(models.Key("host")).(string)
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, paswd, host)

	//Estavblecemos las opciones de conexion del cliente mongo a traves de la url de conexion
	var clientOptions = options.Client().ApplyURI(connectionString)

	//Definimos otras variable spara cachar algun error al conectar
	client, err := mongo.Connect(ctx, clientOptions)
	if err!=nil {
		fmt.Println(err.Error())
		return err
	}

	//Si paso el filtro anterior ahora probamos con un ping al cliente mongo donde sea que este
	err = client.Ping(ctx, nil)
	if err!=nil {
		fmt.Println(err.Error())
		return err
	}

	//Si llego aqui es porque todo salio bien
	fmt.Println("Conexión exitosa con la BD")

	MongoClient = client//Pasamos el valor a la variable global
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

func BaseConectada() bool{
	err:= MongoClient.Ping(context.TODO(),nil)
	return err == nil
}