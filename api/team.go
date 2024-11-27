package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nam-rgba/blv/sqlc"
)

// Get teams request
type GetTeamsRequest struct {
	CoachId int32 `json:"coach_id" binding:"required"`
}

func (s *Server) getTeams(c *gin.Context) {
	coachIdStr := c.Query("coach_id")
	coachId, err := strconv.Atoi(coachIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coach_id"})
		return
	}

	teams, err := s.store.GetTeam(c, sql.NullInt64{Int64: int64(coachId), Valid: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get teams"})
		return
	}
	c.JSON(http.StatusOK, teams)

}

// Create a new team request
type CreateTeamRequest struct {
	Teamname string `json:"teamname" binding:"required"`
	CoachId  int32  `json:"coach_id" binding:"required"`
}

func (s *Server) createTeam(c *gin.Context) {
	var req CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	team, err := s.store.CreateTeam(c, sqlc.CreateTeamParams{
		Teamname: sql.NullString{String: req.Teamname, Valid: true},
		CoachID:  sql.NullInt64{Int64: int64(req.CoachId), Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"team": team,
	})
}

func (s *Server) getTeam(c *gin.Context) {
	teamIdStr := c.Param("team_id")
	teamId, err := strconv.Atoi(teamIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team_id"})
		return
	}
	team, err := s.store.GetTeamById(c, int64(teamId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get team"})
		return
	}
	// Get candidates in the team
	candidates, err := s.store.GetCandidates(c, int64(teamId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get candidates"})
		return
	}
	response := gin.H{
		"teamname":   team.Teamname,
		"coach_id":   team.CoachID,
		"candidates": candidates,
	}
	c.JSON(http.StatusOK, response)
}
