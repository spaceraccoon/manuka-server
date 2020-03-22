package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spaceraccoon/manuka-server/config"
	"github.com/spaceraccoon/manuka-server/models"
	"github.com/spaceraccoon/manuka-server/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("sqlite3", "./database.db") // Need to add env goodness soon

	if err != nil {
		panic(err) // Add better error handling
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Campaign{}, &models.Action{})

	r := routes.SetupRouter()
	r.Run(":3001") // Remember to dotenv this soon
}
