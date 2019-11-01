package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/crouse/pmsPms/internal"
)

// Auth mock
type Auth struct {
	UserFn func(*gin.Context) *model.AuthUser
}

// User mock
func (a *Auth) User(c *gin.Context) *model.AuthUser {
	return a.UserFn(c)
}
