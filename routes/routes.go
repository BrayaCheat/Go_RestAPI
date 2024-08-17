package routes

import (
	"gin-gorm-mysql/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/customer", func(c *gin.Context) {
		controller.GetCustomers(c, db)
	})
	r.GET("/customer/:id", func(c *gin.Context) {
		controller.GetCustomer(c, db)
	})
	r.POST("/customer", func(c *gin.Context) {
		controller.CreateCustomer(c, db)
	})
	r.PUT("/customer/:id", func(c *gin.Context) {
		controller.UpdateCustomer(c, db)
	})
	r.DELETE("/customer/:id", func(c *gin.Context) {
		controller.DeleteUser(c, db)
	})
}
