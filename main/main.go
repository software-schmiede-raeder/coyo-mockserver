package main

import (
	"../mock-service/db"
	"../router"
	cg "./config"
	"github.com/gin-gonic/gin"
	"log"
)

var ginRest *gin.Engine = nil

func main() {
	db.CreateMYSQL(cg.DB_USER, cg.DB_PW, cg.DB_HOST, cg.DB_PORT)
	db.FillDBWithMockData(db.GetDB(cg.DB_USER, cg.DB_PW, cg.DB_HOST, cg.DB_PORT))
	StartGin(cg.GIN_PORT)
}

func StartGin(port string) {
	log.Printf("Starting REST-Service listening on port %s", port)
	if ginRest == nil {
		ginRest = gin.Default()
	}
	router.AddRoutes(ginRest)
	_ = ginRest.Run(":" + port)
}
