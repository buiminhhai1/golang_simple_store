package api

import (
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("*")
	router.Use(cors.New(config))

	router.POST("/auth/register", server.createUser)
	router.POST("/auth/login", server.verifyUser)

	router.POST("/categories", AuthenticateMiddleware, server.createBookCategory)
	router.GET("/categories", AuthenticateMiddleware, server.listCategory)

	router.GET("/books", AuthenticateMiddleware, server.listBook)
	router.POST("/books", AuthenticateMiddleware, server.createBook)

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", AuthenticateMiddleware, server.listUser)
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
