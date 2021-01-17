package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/thuhangpham/simplestores/db/sqlc"
)

// Server serves HTTP requests for our store service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		fmt.Print("hello world go to middleware")
	}, server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
