package main

import (
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spaceraccoon/manuka-server/config"
	"github.com/spaceraccoon/manuka-server/models"
	"github.com/spaceraccoon/manuka-server/routes"
)

var err error

func main() {
	postgresPassword, err := ioutil.ReadFile("/run/secrets/postgres_password")

	if err != nil {
		log.Fatal(err)
	}

	config.DB, err = gorm.Open("postgres", "host=postgres port=5432 sslmode=disable user=postgres dbname=postgres password="+string(postgresPassword))

	if err != nil {
		log.Fatal(err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Campaign{}, &models.Action{})

	r := routes.SetupRouter()
	r.Run()
}
