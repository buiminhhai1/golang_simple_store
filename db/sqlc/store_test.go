package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thuhangpham/simplestores/util"
)

func TestOrderItemTx(t *testing.T) {
	store := NewStore(testDB)
	var books [5]Book
	order := createRandomOrder(t)

	// run n concurrent order item transaction
	n := 5
	errs := make(chan error)
	results := make(chan OrderItemTxResult)

	for i := 0; i < n; i++ {
		book := createRandomBook(t)
		quantity := int32(util.RandomInt(1, int64(book.Quantity)))
		books[i] = book

		go func() {
			result, err := store.OrderItemTx(context.Background(), OrderItemTxParams{
				OrderID:      order.ID,
				BookID:       book.ID,
				BookQuantity: book.Quantity,
				BookPrice:    book.Price,
				Quantity:     quantity,
			})
			errs <- err
			results <- result
		}()
	}

	// check result
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check orderItem
		orderItem := result.OrderItem
		require.NotEmpty(t, orderItem)
		require.Equal(t, orderItem.OrderID, order.ID)
		require.NotZero(t, orderItem.BookID)
		require.NotZero(t, orderItem.PriceInvoice)
		require.NotZero(t, orderItem.Quantity)
		require.NotZero(t, orderItem.CreatedAt)

		// check book
		teBook := result.Book
		require.NotEmpty(t, teBook)
		for _, book := range books {
			if book.ID == teBook.ID {
				require.Greater(t, book.Quantity, teBook.Quantity)
			}
		}
		// check orderitem
		teOrderItem := result.OrderItem

		_, err = store.GetOrderItemByPrimaryKey(context.Background(), GetOrderItemByPrimaryKeyParams{OrderID: teOrderItem.OrderID, BookID: teOrderItem.BookID})
		require.NoError(t, err)
	}
}
