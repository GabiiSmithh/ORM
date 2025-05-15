package models

type Pessoa struct {
    Nome      string  `bson:"nome"`
    DataNasc  string  `bson:"data_nasc"`
    CPF       string  `bson:"cpf"`
    Telefone  string  `bson:"telefone"`
    Altura    float64 `bson:"altura"`
}
