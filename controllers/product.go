package controllers

import (
	"net/http"
	"photo-app/database"
	"photo-app/helpers"
	"photo-app/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Products:= models.Product{}
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Products)
	} else {
		c.ShouldBind(&Products)
	}
	Products.UserID = userID

	Result := map[string]interface{}{}

	err := db.Raw(
		"INSERT into products (title, caption, photo_url, harga, user_id, created_at) VALUES(?,?,?,?,?,?) Returning id,title,photo_url, harga, user_id, created_at, caption",
		Products.Title, Products.Caption, Products.PhotoUrl, Products.Harga, Products.UserID, time.Now(),
	).Scan(&Result).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Result)

}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Products := models.Product{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Products)
	} else {
		c.ShouldBind(&Products)
	}
	Products.UserID = userID

	Result := map[string]interface{}{}
	SqlStatement := "Update products SET title = ?, caption = ?, photo_url = ?, harga = ?, updated_at = ? WHERE id = ? RETURNING id, title, caption, photo_url, harga, updated_at"
	err := db.Raw(
		SqlStatement,
		Products.Title, Products.Caption, Products.PhotoUrl, Products.Harga, time.Now(), uint(photoId),
	).Scan(&Result).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Result)
}

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	// data all yang bener
	Products := []models.Product{}
	err := db.Find(&Products).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Products)
}

func GetOneProduct(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Products := models.Product{}
	err := db.Preload("User").Find(&Products, uint(photoId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Products)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Products := models.Product{}
	err := db.Delete(Products, uint(photoId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
