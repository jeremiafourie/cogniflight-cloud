package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/jeremiafourie/cogniflight-cloud/backend/types"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func HashPwd(pwd string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPwd(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
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

		if CheckPwd(user.Pwd, req.Pwd) {
			c.Status(200)
			sess, _ := s.CreateSession(user.ID, user.Role)
			c.SetCookie("sessid", sess.SessID, 3600, "/", "", false, true)
		} else {
			c.Status(401)
		}
	}
}
