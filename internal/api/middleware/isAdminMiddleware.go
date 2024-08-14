package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IsAdminMiddleWare struct{}

func NewIsAdminMiddleWare() *IsAdminMiddleWare {
	return &IsAdminMiddleWare{}
}

func (m *IsAdminMiddleWare) VerificarIsAdminMiddleWare(c *gin.Context) {
	isAdmin, exists := c.Get("IsAdmin")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Key not exists"})
		c.Abort()
	}
	if isAdmin == false {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unauthorized"})
		c.Abort()
	}
	c.Next()
}
