package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nam-rgba/blv/auth"
	"github.com/nam-rgba/blv/sqlc"
)

type CreateUserRequest struct {
	Role     string `json:"role" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Role     string `json:"role" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Signup creates a new user
func (s *Server) Signup(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// If role is candidate, register candidate
	if req.Role == "candidate" {

		_, err := s.store.RegisterCandidate(ctx, sqlc.RegisterCandidateParams{
			Email:    sql.NullString{String: req.Email, Valid: true},
			Password: hashedPassword,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"status": "Candidate created"})
	}

	// If role is coach, register coach
	if req.Role == "coach" {
		_, err := s.store.RegisterCoach(ctx, sqlc.RegisterCoachParams{
			Email:    sql.NullString{String: req.Email, Valid: true},
			Password: hashedPassword,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"status": "Coach created"})
	}
}

// Login authenticates a user
func (s *Server) Login(c *gin.Context) {
	var req LoginUserRequest

	// Check is request is valid
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check role
	if req.Role == "candidate" {
		candidate, err := s.store.GetCandidateByEmail(c, sql.NullString{String: req.Email, Valid: true})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Check hash password error
		errHash := auth.ComparePassword(candidate.Password, req.Password)
		if errHash != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errHash.Error()})
			return
		}

		// Generate token
		token, err := auth.GenerateToken(
			candidate.Email.String,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"token":     token,
			"role":      "candidate",
			"user_id":   candidate.CanID,
			"user_mail": candidate.Email.String,
			"success":   "user logged in",
		})
		return

	} else if req.Role == "coach" {
		coach, err := s.store.GetCoachByEmail(c, sql.NullString{String: req.Email, Valid: true})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check hash password error
		errHash := auth.ComparePassword(coach.Password, req.Password)
		if errHash != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errHash.Error()})
			return
		}

		// Generate token
		token, err := auth.GenerateToken(
			coach.Email.String,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		response := gin.H{
			"token":     token,
			"role":      "coach",
			"user_id":   coach.CoachID,
			"user_mail": coach.Email.String,
			"success":   "user logged in",
		}
		fmt.Println(response)
		c.JSON(200, gin.H{
			"token":     token,
			"role":      "coach",
			"user_id":   coach.CoachID,
			"user_mail": coach.Email.String,
			"success":   "user logged in",
		})
		return
	}

}
