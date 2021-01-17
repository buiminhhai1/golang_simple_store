package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thuhangpham/simplestores/util"
)

func createRandomOrderItem(t *testing.T) OrderItem {
	order := createRandomOrder(t)
	book := createRandomBook(t)

	arg := CreateOrderItemParams{
		OrderID:      order.ID,
		BookID:       book.ID,
		Quantity:     int32(util.RandomInt(1, int64(book.Quantity))),
		PriceInvoice: book.Price,
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.BookID, orderItem.BookID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, arg.PriceInvoice, orderItem.PriceInvoice)
	require.NotZero(t, orderItem.CreatedAt)

	return orderItem
}

func createRandomOrderWithOrderID(t *testing.T, orderID int32) OrderItem {
	book := createRandomBook(t)

	arg := CreateOrderItemParams{
		OrderID:      orderID,
		BookID:       book.ID,
		Quantity:     int32(util.RandomInt(1, int64(book.Quantity))),
		PriceInvoice: book.Price,
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.BookID, orderItem.BookID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, arg.PriceInvoice, orderItem.PriceInvoice)
	require.NotZero(t, orderItem.CreatedAt)

	return orderItem
}

func createRandomOrderWithBookInfo(t *testing.T, bookID int32, bookQuantity int32, bookPrice int32) OrderItem {
	order := createRandomOrder(t)

	arg := CreateOrderItemParams{
		OrderID:      order.ID,
		BookID:       bookID,
		Quantity:     int32(util.RandomInt(1, int64(bookQuantity))),
		PriceInvoice: bookPrice,
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.BookID, orderItem.BookID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, arg.PriceInvoice, orderItem.PriceInvoice)
	require.NotZero(t, orderItem.CreatedAt)

	return orderItem
}

func TestCreateOrderItem(t *testing.T) {
	orderItem1 := createRandomOrderItem(t)

	arg := GetOrderItemByPrimaryKeyParams{
		OrderID: orderItem1.OrderID,
		BookID:  orderItem1.BookID,
	}

	orderItem2, err := testQueries.GetOrderItemByPrimaryKey(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orderItem2)

	require.Equal(t, orderItem1.OrderID, orderItem2.OrderID)
	require.Equal(t, orderItem1.BookID, orderItem2.BookID)
	require.Equal(t, orderItem1.Quantity, orderItem2.Quantity)
	require.Equal(t, orderItem1.PriceInvoice, orderItem2.PriceInvoice)
	require.WithinDuration(t, orderItem1.CreatedAt, orderItem2.CreatedAt, time.Second)
}

func TestUpdateOrderPriceInvoice(t *testing.T) {
	orderItem1 := createRandomOrderItem(t)

	arg := UpdateOrderPriceInvoiceParams{
		OrderID:      orderItem1.OrderID,
		BookID:       orderItem1.BookID,
		PriceInvoice: int32(util.RandomInt(400, 2000)),
	}

	orderItem2, err := testQueries.UpdateOrderPriceInvoice(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orderItem2)

	require.Equal(t, orderItem2.OrderID, orderItem2.OrderID)
	require.Equal(t, arg.BookID, orderItem2.BookID)
	require.Equal(t, arg.PriceInvoice, orderItem2.PriceInvoice)
	require.Equal(t, orderItem1.Quantity, orderItem2.Quantity)

	require.WithinDuration(t, orderItem2.CreatedAt, orderItem2.CreatedAt, time.Second)
}

func TestDeleteOrderItem(t *testing.T) {
	orderItem1 := createRandomOrderItem(t)

	arg := DeleteOrderItemParams{
		OrderID: orderItem1.OrderID,
		BookID:  orderItem1.BookID,
	}

	err := testQueries.DeleteOrderItem(context.Background(), arg)
	require.NoError(t, err)

	argGet := GetOrderItemByPrimaryKeyParams{
		OrderID: orderItem1.OrderID,
		BookID:  orderItem1.BookID,
	}
	orderItem2, err := testQueries.GetOrderItemByPrimaryKey(context.Background(), argGet)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, orderItem2)
}

func TestGetOrderItemsByOrderID(t *testing.T) {
	order := createRandomOrder(t)

	for i := 0; i < 10; i++ {
		createRandomOrderWithOrderID(t, order.ID)
	}

	arg := GetOrderItemsByOrderIDParams{
		Limit:   5,
		Offset:  5,
		OrderID: order.ID,
	}

	orderItems, err := testQueries.GetOrderItemsByOrderID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orderItems, 5)

	for _, item := range orderItems {
		require.NotEmpty(t, item)
	}
}

func TestGetOrderItemsByBookID(t *testing.T) {
	book := createRandomBook(t)

	for i := 0; i < 10; i++ {
		createRandomOrderWithBookInfo(t, book.ID, book.Quantity, book.Price)
	}

	arg := GetOrderItemsByBookIDParams{
		Limit:  5,
		Offset: 5,
		BookID: book.ID,
	}

	orderItems, err := testQueries.GetOrderItemsByBookID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orderItems, 5)

	for _, item := range orderItems {
		require.NotEmpty(t, item)
	}
}
