package controllers

import (
	"net/http"
	"photo-app/database"
	"photo-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)


func UpdateItemOrder(c *gin.Context){
	db := database.GetDB()

	// Check data exist
	var orders models.Order

	err := db.First(&orders, "id = ?", c.Param("orderid")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	

	// Input Update Order
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}), // column needed to be updated
	}).Create(&input)

	// Input Update Item
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"item_code", "description", "quantity"}), // column needed to be updated
	}).Create(&input.Items)


	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}	