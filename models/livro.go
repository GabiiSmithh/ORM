package models

type Livro struct {
	ID        string `bson:"_id"` 
	ISBN      string `bson:"isbn"`   
	Titulo    string `bson:"titulo"`
	Autor     string `bson:"autor"`
	AnoPublic int    `bson:"ano_public"`
}
