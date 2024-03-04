package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snoopy910/golang-gorm-postgres/controllers"
	"github.com/snoopy910/golang-gorm-postgres/initializers"
)

var (
	server *gin.Engine
	StationController controllers.StationController
	// IndegoRouteController routes.IndegoRouteController

)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db := initializers.ConnectDB(&config)

	StationController = controllers.NewStationController(db)

	// controllers.CreateData(db, indego)

	// IndegoController = controllers.NewIndegoController(db)
	// IndegoRouteController = routes.NewRouteIndegoController(IndegoController)
	
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not lod environment variables", err)
	}

	router := server.Group("/api/"+config.Version+"/stations")
	router.GET("healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	router.POST("/data-fetch-and-store-it-db", StationController.DataFetchAndStore)
	router.GET("", StationController.GetByTime)
	router.GET("/:kioskId", StationController.GetByKiosk)

	// IndegoRouteController.IndegoRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}