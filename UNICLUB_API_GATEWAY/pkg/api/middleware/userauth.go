package middleware

import (
	"fmt"
	"net/http"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var cfg config.Config

func CfgHelper(conf config.Config) {

	cfg = conf
}

func UserAuthMiddleware(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
  
  fmt.Println("sdklfjsdljfsdkjflsdj\n\n\n\n\n",cfg.USERACCESSSECRET)

	// tokenString = strings.TrimPrefix(tokenString, "Bearer ")
  
  fmt.Println("token",tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
  
		return []byte(cfg.USERACCESSSECRET), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user authorization token"})
		c.Abort()
		return
	}
	// var userID models.TokenUsersID
	// uID := claims["id"].(float64)

	// userID.UserID=uint(uID)

	fmt.Println("claims", claims)

	role, ok := claims["role"].(string)
	if !ok || role != "client" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "error in retrieving id"})
		c.Abort()
		return
	}

	c.Set("role", role)
	c.Set("id", int(id))

	c.Next()
}
