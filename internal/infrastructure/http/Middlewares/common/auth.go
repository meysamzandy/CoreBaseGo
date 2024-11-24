package middlewares

import (
	"CoreBaseGo/internal/interfaces/rest"
	messages "CoreBaseGo/internal/interfaces/rest/Messages"
	"CoreBaseGo/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(tokenKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			rest.JSONOutput(c, http.StatusUnauthorized, nil, messages.MissedToken, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)
		claims, err := utils.VerifyJWT(tokenKey, tokenString)
		if err != nil {
			fmt.Println(err.Error())
			rest.JSONOutput(c, http.StatusUnauthorized, nil, messages.InvalidToken, "Invalid token :"+err.Error())
			c.Abort()
			return
		}

		// If token is valid, set the user to the context
		c.Set("user", claims)
		c.Next()
	}
}
