package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error){
	costo:= 8//Las vueltas que va a ir a encriptar el password y comparar si esta bien, por seguridad
	//6= aceptable, 8=Muy bien, 10= Muy costosa y mas segura
	bytes, err:= bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err!=nil {
		return err.Error(), err
	}

	return string(bytes), nil
}