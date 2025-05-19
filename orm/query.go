package orm

import (
	"go.mongodb.org/mongo-driver/bson"
)

//agrupa os parâmetros opcionais para consultas customizadas
type QueryOptions struct {
	Filter     interface{}   // filtro de consulta (ex: bson.M{"nome": "João"})
	Projection bson.D        // campos a retornar (ex: bson.D{{"nome", 1}, {"email", 1}})
	Sort       bson.D        // ordenação (ex: bson.D{{"nome", 1}})
}

