package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thuhangpham/simplestores/util"
)

func createRandomOrder(t *testing.T) Order {
	user := createRandomUser(t)

	arg := CreateOrderParams{
		UserID: user.ID,
		Status: OderenumPending,
	}
	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, user.ID, order.UserID)
	require.Equal(t, arg.Status, order.Status)
	require.NotZero(t, order.CreatedAt)

	return order
}

func createRandomOrderWithUserID(t *testing.T, userID int32) Order {
	arg := CreateOrderParams{
		UserID: userID,
		Status: OderenumPending,
	}
	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserID, order.UserID)
	require.Equal(t, arg.Status, order.Status)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestCreateOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	order2, err := testQueries.GetOrderById(context.Background(), order1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.UserID, order2.UserID)
	require.Equal(t, order1.Status, order2.Status)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestUpdateOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	arg := UpdateOrderStatusParams{
		ID:     order1.ID,
		Status: Oderenum(util.RandomOrderStatus()),
	}
	order2, err := testQueries.UpdateOrderStatus(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, arg.Status, order2.Status)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestDeleteOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	err := testQueries.DeleteOrderById(context.Background(), order1.ID)
	require.NoError(t, err)

	order2, err := testQueries.GetOrderById(context.Background(), order1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, order2)
}

func TestGetOrdersByUserId(t *testing.T) {
	user := createRandomUser(t)

	for i := 0; i < 10; i++ {
		createRandomOrderWithUserID(t, user.ID)
	}

	arg := GetOrdersByUserIdParams{
		Limit:  5,
		Offset: 5,
		UserID: user.ID,
	}

	orders, err := testQueries.GetOrdersByUserId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}

func TestListOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	orders, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}
