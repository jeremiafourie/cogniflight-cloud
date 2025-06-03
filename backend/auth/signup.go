package auth

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jeremiafourie/cogniflight-cloud/backend/types"
)

func CreateSignupToken(s types.SignupTokenStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email string     `json:"email"`
			Phone string     `json:"phone"`
			Role  types.Role `json:"role" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Email == "" && req.Phone == "" {
			c.JSON(400, gin.H{"error": "Please provide either email, or phone"})
			return
		}

		tok, err := s.CreateSignupToken(req.Phone, req.Email, req.Role, 6*time.Hour)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal error"})
			log.Printf("Error creating signup token: %v", err)
			return
		}

		c.JSON(201, gin.H{"tokStr": tok.TokStr})
	}
}
