package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db := initDB()
	// defer db.Close()

	productController := InitProductController(nil)

	r := gin.Default()

	r.GET("/products", productController.FindAll)
	r.GET("/products/:id", productController.FindByID)
	r.POST("/products", productController.Create)
	r.PUT("/products/:id", productController.Update)
	r.DELETE("/products/:id", productController.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
