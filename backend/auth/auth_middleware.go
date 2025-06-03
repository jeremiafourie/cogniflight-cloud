package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jeremiafourie/cogniflight-cloud/backend/types"
)

func AuthMiddleware(s types.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess_id, err := c.Cookie("sessid")
		if err != nil {
			c.AbortWithStatus(401)
		}

		_, err = s.GetSession(sess_id)
		if err != nil {
			c.AbortWithStatus(401)
		}
	}
}
