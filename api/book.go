package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/thuhangpham/simplestores/db/sqlc"
)

type createBookRequest struct {
	Name           string `json:"name" binding:"required"`
	Price          int32  `json:"price" binding:"required,min=1"`
	Quantity       int32  `json:"quantity" binding:"required,min=1"`
	Image          string `json:"image" binding:"required"`
	Description    string `json:"description" binding:"required"`
	BookCategoryID int32  `json:"book_category_id" binding:"required"`
}

func (server *Server) createBook(ctx *gin.Context) {
	var req createBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBookParams{
		Name:           req.Name,
		Price:          req.Price,
		Quantity:       req.Quantity,
		Image:          req.Image,
		Description:    req.Description,
		BookCategoryID: req.BookCategoryID,
	}
	book, err := server.store.CreateBook(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type listBookRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listBook(ctx *gin.Context) {
	var req listBookRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBooksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	books, err := server.store.ListBooks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, books)
}
