package main

import (
	"github.com/gin-gonic/gin"
	"github.com/macadadi/bookstore/dbhandler"
	"github.com/macadadi/bookstore/services"
)

func main() {
	db := dbhandler.DBconnect()
	server := gin.Default()
	r := server.Group("v1")

	r.GET("/users", services.GetUsers(db))
	r.GET("/user/:id", services.GetSingleUser(db))
	r.POST("/user", services.AddNewUser(db))
	r.PUT("/user/:id", services.UpdateUser(db))
	r.DELETE("/user/:id", services.Delete(db))

	server.Run(":8080")
}
