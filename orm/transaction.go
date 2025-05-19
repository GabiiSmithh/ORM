package orm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//define uma operação que aceita um contexto de sessão
type TxOperation func(sessCtx mongo.SessionContext) error

// executa uma lista de operações dentro de uma transação
func RunTransaction(client *mongo.Client, ops []TxOperation) error {
	session, err := client.StartSession()
	if err != nil {
		return fmt.Errorf("falha ao iniciar sessão: %w", err)
	}
	defer session.EndSession(context.Background())

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		for _, op := range ops {
			if err := op(sessCtx); err != nil {
				return nil, fmt.Errorf("erro na operação: %w", err)
			}
		}
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return fmt.Errorf("falha na transação: %w", err)
	}
	return nil
}
