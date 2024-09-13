package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) Logout(c *gin.Context) {
	defer c.Redirect(http.StatusFound, "/login")
	c.SetCookie("access-token", "", -1, "/", "", false, true)
}
