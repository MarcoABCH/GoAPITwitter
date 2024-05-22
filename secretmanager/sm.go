package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/MarcoABCH/GoAPITwitter/awsgo"
	"github.com/MarcoABCH/GoAPITwitter/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error){
	var datosSecret models.Secret

	fmt.Println("> Pido Secret " + secretName)//Estos mensajes van quedando en Cloudwatch AWS

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err!=nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)//Estamos usanod punteros para ir a guardar los datos Secret
	fmt.Println("> Lectura de Secret OK " + secretName)

	return datosSecret, nil
}