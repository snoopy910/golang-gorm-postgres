package main

import (
	"fmt"
	"log"

	"github.com/snoopy910/golang-gorm-postgres/initializers"
	"github.com/snoopy910/golang-gorm-postgres/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environemnt variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Debug().AutoMigrate(&models.Core{})
	fmt.Println("? Migration Complete")
}