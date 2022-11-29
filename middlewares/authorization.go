package middlewares

import (
	// "fmt"
	"net/http"
	"photo-app/database"
	"photo-app/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		// Data User Define
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}
		// solving case 1
		userLevel := db.Select("level").First(&User, "id = ?", uint(userID))
		_ = userLevel
		if c.Request.Method != "POST" {
			photoId, err := strconv.Atoi(c.Param("photoId"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "invalid parameter",
				})
				return
			}

			// Data Product Define
			Products := models.Product{}
			err = db.Select("user_id").First(&Products, uint(photoId)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doesn't exist",
				})
				return
			}

				// User Level Authorization
			if User.Level == "admin" || User.Level == "superadmin"{
				return
			} else if User.Level == "user" {
				// Data Photo Authorization
				if Products.UserID != userID {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error":   "Unauthorized",
						"message": "you are not allowed to access this data",
					})
					return
				}
			}else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}else if c.Request.Method == "POST" {
			if User.Level == "admin" || User.Level == "superadmin"{
				return
			} else if User.Level == "user" {
				// Data Photo Authorization
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}

		c.Next()
	}
}
