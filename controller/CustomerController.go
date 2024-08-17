package controller

import (
	"gin-gorm-mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCustomers(c *gin.Context, db *gorm.DB) {
	var customers []models.Customer
	db.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func GetCustomer(c *gin.Context, db *gorm.DB) {
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	db.Find(&customer)
	c.JSON(http.StatusOK, customer)
}

func CreateCustomer(c *gin.Context, db *gorm.DB) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	db.Create(&customer)
	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(c *gin.Context, db *gorm.DB) {
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&customer)
	c.Status(http.StatusNoContent)
}
