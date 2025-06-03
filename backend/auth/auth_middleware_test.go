package auth

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware(t *testing.T) {
	sessionStore := &FakeSessionStore{}

	r := gin.New()

}
