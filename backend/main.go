package main

import (
	"github.com/nitaxxix/sa-64-example/controller"

	"github.com/nitaxxix/sa-64-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// RemedyType

	r.GET("/remedy_types", controller.ListRemedyType)

	r.POST("/remedy_type", controller.CreateRemedyType)

	// Patient

	r.GET("/patients", controller.ListPatient)

	r.POST("/patient", controller.CreatePatient)

	// Appoint

	r.GET("/appoints", controller.ListAppoint)

	r.POST("/appoint", controller.CreateAppoint)

	// user

	r.GET("/users", controller.ListUser)

	r.GET("/users/:id", controller.GetUser)

	r.POST("/user", controller.CreateUser)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
