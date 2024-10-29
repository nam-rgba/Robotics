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

	router.POST("/signup", server.Signup)

	router.POST("/login/candidate", server.LoginCandidate)
	router.POST("/login/coach", server.LoginCoach)

	// add routes to router
	server.router = router

	return server
}

// Start run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
