package models

type Produto struct {
	ID         string  `bson:"_id"`
	Nome       string  `bson:"nome"`
	Codigo     string     `bson:"codigo"`
	Categoria  string  `bson:"categoria"`
	Preco      float64 `bson:"preco"`
	Quantidade int     `bson:"quantidade"`
}
