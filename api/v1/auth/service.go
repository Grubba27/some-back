package auth

import (
	"app/lib/cripto"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		key := c.Request.Header.Get("Authorization")

		if key == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		claims := &cripto.Claims{}

		token, err := jwt.ParseWithClaims(key, claims, func(token *jwt.Token) (interface{}, error) {
			return cripto.Key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
				return
			}

			c.AbortWithStatusJSON(400, gin.H{"message": "Bad Request", "description": err.Error()})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("userID", claims.ID)

		c.Next()
	}
}
