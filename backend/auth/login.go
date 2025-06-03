package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/jeremiafourie/cogniflight-cloud/backend/types"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func Login(u types.UserStore, s types.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email string `json:"email"`
			Pwd   string `json:"pwd"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.Status(400)
			return
		}

		user, err := u.GetUserByEmail(req.Email)
		if err != nil {
			c.Status(401)
			return
		}

		if req.Pwd == user.Pwd {
			c.Status(200)
			sess, _ := s.CreateSession(user.ID, user.Role)
			c.SetCookie("sessid", sess.SessID, 3600, "/", "", false, true)
		} else {
			c.Status(401)
		}
	}
}
