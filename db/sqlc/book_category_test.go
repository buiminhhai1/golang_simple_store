package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thuhangpham/simplestores/util"
)

func createRandomCategory(t *testing.T) BookCategory {
	name := util.RandomString(12)
	category, err := testQueries.CreateCategory(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, name, category.Name)

	require.NotZero(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	cate1 := createRandomCategory(t)

	cate2, err := testQueries.GetCategory(context.Background(), cate1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, cate2)

	require.Equal(t, cate1.Name, cate2.Name)
	require.WithinDuration(t, cate1.CreatedAt, cate2.CreatedAt, time.Second)
}

func TestUpdateCategory(t *testing.T) {
	cate1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:   cate1.ID,
		Name: util.RandomString(10),
	}

	cate2, err := testQueries.UpdateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, cate2)

	require.Equal(t, arg.Name, cate2.Name)
	require.WithinDuration(t, cate1.CreatedAt, cate2.CreatedAt, time.Second)
}

func TestDeleteCategory(t *testing.T) {
	cate1 := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), cate1.ID)

	require.NoError(t, err)

	cate2, err := testQueries.GetCategory(context.Background(), cate1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, cate2)
}

func TestGetCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}

	arg := ListBookCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	categories, err := testQueries.ListBookCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}

}
