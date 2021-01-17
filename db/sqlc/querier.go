// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateCategory(ctx context.Context, name string) (BookCategory, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteBook(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	DeleteOrderById(ctx context.Context, id int32) error
	DeleteOrderItem(ctx context.Context, arg DeleteOrderItemParams) error
	DeleteUser(ctx context.Context, id int32) error
	GetBookById(ctx context.Context, id int32) (Book, error)
	GetBooksByCategoryId(ctx context.Context, arg GetBooksByCategoryIdParams) ([]Book, error)
	GetCategory(ctx context.Context, id int32) (BookCategory, error)
	GetOrderById(ctx context.Context, id int32) (Order, error)
	GetOrderItemByPrimaryKey(ctx context.Context, arg GetOrderItemByPrimaryKeyParams) (OrderItem, error)
	GetOrderItemsByBookID(ctx context.Context, arg GetOrderItemsByBookIDParams) ([]OrderItem, error)
	GetOrderItemsByOrderID(ctx context.Context, arg GetOrderItemsByOrderIDParams) ([]OrderItem, error)
	GetOrdersByUserId(ctx context.Context, arg GetOrdersByUserIdParams) ([]Order, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByUserName(ctx context.Context, username string) (User, error)
	ListBookCategories(ctx context.Context, arg ListBookCategoriesParams) ([]BookCategory, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (BookCategory, error)
	UpdateOrderPriceInvoice(ctx context.Context, arg UpdateOrderPriceInvoiceParams) (OrderItem, error)
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (Order, error)
	UpdateQuantityBook(ctx context.Context, arg UpdateQuantityBookParams) (Book, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
