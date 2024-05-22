package models

type Secret struct {
	Host     string `json:"host"` //alt + 96 para este carcater
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSign  string `json:"jwtsign"`
	DataBase string `json:"database"`
}