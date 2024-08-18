package main

import (
	"gin-gorm-mysql/config"
	"gin-gorm-mysql/middleware"
	"gin-gorm-mysql/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the database
	db := config.InitDB()

	//set up gin router
	r := gin.Default()

	// apply the logger middleware globally
	r.Use(middleware.Logger())

	//register routes
	routes.RegisterRoutes(r, db)

	//run the server
	r.Run(":8081")
}
