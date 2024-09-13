package router

import (
	"net/http"

	"github.com/chscz/backend-onboard-task/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(ua *auth.UserAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		if (c.Request.URL.Path == "/login") || (c.Request.URL.Path == "/register") {
			c.Next()
			return
		}

		accessToken, err := c.Cookie("access-token")
		if err != nil {
			c.Abort()
			c.Redirect(http.StatusFound, "/login")
			return
		}

		claims, err := ua.ValidateJWT(ua.JWTSecretKey, accessToken)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
				c.SetCookie("access-token", "", -1, "/", "", false, true)
			}
			c.Abort()
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
