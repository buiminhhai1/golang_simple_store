package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transactions
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// OrderItemTxParams contains the input parameters of the order item transaction
type OrderItemTxParams struct {
	OrderID      int32 `json:"order_id"`
	BookID       int32 `json:"book_id"`
	BookQuantity int32 `json:"book_quantity"`
	BookPrice    int32 `json:"book_price"`
	Quantity     int32 `json:"quantity"`
}

// OrderItemTxResult is the result of the order item transction
type OrderItemTxResult struct {
	OrderItem OrderItem `json:"order_item"`
	Book      Book      `json:"book"`
}

// OrderItemTx performs a quantity of book to order by use
// It creates a orderItem records, and update book within a single database transaction
func (store *Store) OrderItemTx(ctx context.Context, arg OrderItemTxParams) (OrderItemTxResult, error) {
	var result OrderItemTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.OrderItem, err = q.CreateOrderItem(ctx, CreateOrderItemParams{
			OrderID:      arg.OrderID,
			BookID:       arg.BookID,
			Quantity:     arg.Quantity,
			PriceInvoice: arg.BookPrice,
		})

		if err != nil {
			return err
		}

		result.Book, err = q.UpdateQuantityBook(ctx, UpdateQuantityBookParams{
			ID:       arg.BookID,
			Quantity: arg.BookQuantity - arg.Quantity,
		})

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
