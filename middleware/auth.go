package middleware

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	var expiredAt interface{}
	for key, val := range claims {
		if key == "expiredAt" {
			expiredAt = val
		}
	}
	if time.Now().Unix() > int64(expiredAt.(float64)) {
		result := gin.H{
			"message": "not authorized",
			"status":  "token expired",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
	if token != nil && err == nil {
	} else {
		result := gin.H{
			"message": "not authorized",
			"status":  err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
