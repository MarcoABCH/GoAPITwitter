package bd

import (
	"context"

	"github.com/MarcoABCH/GoAPITwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string){
	ctx:= context.TODO()

	db := MongoClient.Database(DatabaseName)

	col := db.Collection("usuarios")

	condition := bson.M{"email" : email}

	var resultado models.Usuario

	err:= col.FindOne(ctx, condition).Decode(&resultado)

	ID:= resultado.ID.Hex() //Converite a string el Id para pasarlo y devolverlo
	if err!= nil{
		return resultado, false, ID 
	}

	//si no hay error al final enviamos el ID
	return resultado, true, ID

}