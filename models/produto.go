package models

type Produto struct {
    Nome      string  `bson:"nome"`
    Preco     float64 `bson:"preco"`
    Quantidade int    `bson:"quantidade"`
}