package models

type Pedido struct {
    ClienteID string    `bson:"cliente_id"`
    Produtos  []Produto `bson:"produtos"`
    Data      string    `bson:"data"` // pode ser string ou time.Time, você escolhe
}