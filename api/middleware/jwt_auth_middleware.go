package middleware

import (
	"go_backend_clean_code/domain"
	"go_backend_clean_code/internal/tokenutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// takes the access secret key as an input parameter.
func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// extracts the token from the header of the request
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			//checks if the token is authorized or not, and it returns an error if not authorized.
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				//extracts the UserID from the token
				//that we have put in while creating the access token
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				// put it into the context of the HTTP web framework used
				//so that we can extract it easily
				//when needed later in the request flow.
				c.Set("x-user-id", userID)
				//We can get the UserID from the HTTP Web Framework Context
				//like userID := c.GetString("x-user-id")
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}