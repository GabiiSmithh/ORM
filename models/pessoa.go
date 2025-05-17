package models

type Pessoa struct {
	ID       string  `bson:"_id"`
	Nome     string  `bson:"nome"`
	CPF      string  `bson:"cpf"`
	DataNasc string  `bson:"data_nasc"`
	Telefone string  `bson:"telefone"`
	Altura   float64 `bson:"altura"`
}
