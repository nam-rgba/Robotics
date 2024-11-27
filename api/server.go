package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nam-rgba/blv/sqlc"
)

// Server serves HTTP request
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Create a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(corsMiddleware())

	router.POST("/signup", server.Signup)

	router.POST("/login", server.Login)

	router.GET("/teams", server.getTeams)

	router.POST("/teams/create", server.createTeam)

	router.GET("/teams/:team_id", server.getTeam)

	// add routes to router
	server.router = router

	return server
}

// Start run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Router of the HTTP server
func (server *Server) Router() *gin.Engine {
	return server.router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
