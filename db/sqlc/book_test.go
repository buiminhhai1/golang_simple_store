package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thuhangpham/simplestores/util"
)

func createRandomBook(t *testing.T) Book {
	cate := createRandomCategory(t)
	arg := CreateBookParams{
		Name:           util.RandomString(20),
		Quantity:       int32(util.RandomInt(10, 20)),
		Price:          int32(util.RandomInt(500, 2000)),
		BookCategoryID: cate.ID,
	}

	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	return book
}

func createRandomBookWithCategoryID(t *testing.T, id int32) Book {
	arg := CreateBookParams{
		Name:           util.RandomString(20),
		Quantity:       int32(util.RandomInt(10, 20)),
		Price:          int32(util.RandomInt(500, 2000)),
		BookCategoryID: id,
	}

	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	return book
}

func TestCreateBook(t *testing.T) {
	book1 := createRandomBook(t)

	book2, err := testQueries.GetBookById(context.Background(), book1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.BookCategoryID, book2.BookCategoryID)
	require.Equal(t, book1.Name, book2.Name)
	require.Equal(t, book1.Price, book2.Price)
	require.Equal(t, book1.Quantity, book2.Quantity)
	require.WithinDuration(t, book1.CreatedAt, book2.CreatedAt, time.Second)
}

func TestUpdateBook(t *testing.T) {
	book1 := createRandomBook(t)

	arg := UpdateBookParams{
		ID:             book1.ID,
		Name:           util.RandomString(15),
		Price:          int32(util.RandomInt(500, 5000)),
		Quantity:       int32(util.RandomInt(10, 20)),
		BookCategoryID: book1.BookCategoryID,
	}

	book2, err := testQueries.UpdateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, arg.Name, book2.Name)
	require.Equal(t, arg.Price, book2.Price)
	require.Equal(t, arg.Quantity, book2.Quantity)
	require.Equal(t, arg.BookCategoryID, book2.BookCategoryID)
}

func TestUpdateQuantityBook(t *testing.T) {
	book1 := createRandomBook(t)

	arg := UpdateQuantityBookParams{
		ID:       book1.ID,
		Quantity: int32(int64(book1.Quantity) - (util.RandomInt(1, int64(book1.Quantity)))),
	}

	book2, err := testQueries.UpdateQuantityBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, arg.Quantity, book2.Quantity)
}

func TestDeleteBook(t *testing.T) {
	book1 := createRandomBook(t)

	err := testQueries.DeleteBook(context.Background(), book1.ID)
	require.NoError(t, err)

	book2, err := testQueries.GetBookById(context.Background(), book1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book2)

}

func TestGetBooks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBook(t)
	}

	arg := ListBooksParams{
		Limit:  5,
		Offset: 5,
	}

	books, err := testQueries.ListBooks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, books, 5)

	for _, book := range books {
		require.NotEmpty(t, book)
	}
}

func TestGetBooksByCategoryID(t *testing.T) {
	category := createRandomCategory(t)
	for i := 0; i < 10; i++ {
		createRandomBookWithCategoryID(t, category.ID)
	}

	arg := GetBooksByCategoryIdParams{
		Limit:          5,
		Offset:         5,
		BookCategoryID: category.ID,
	}

	books, err := testQueries.GetBooksByCategoryId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, books, 5)

	for _, book := range books {
		require.NotEmpty(t, book)
	}
}
