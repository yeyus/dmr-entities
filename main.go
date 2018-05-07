package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/yeyus/dmr-entities/brandmeister"
	"github.com/yeyus/dmr-entities/db"
	"github.com/yeyus/dmr-entities/dmrmarc"
	"github.com/yeyus/dmr-entities/model"
	"github.com/yeyus/dmr-entities/utils"
	"log"
	"strconv"
)

type Configuration struct {
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseName     string `envconfig:"DB_NAME"`
	Debug            bool   `envconfig:"DEBUG"`
}

func main() {
	var c Configuration
	err := envconfig.Process("entities", &c)
	if err != nil {
		log.Fatal(err.Error)
	}

	database := db.GetDBManager(c.DatabaseHost, c.DatabaseUser, c.DatabasePassword, c.DatabaseName, c.Debug)
	err = database.DropSchema()
	err = database.CreateSchema()
	if err != nil {
		log.Panicf("Error creating schema: %v", err)
	}

	groups := brandmeister.GetBrandmeisterGroups()
	for k, v := range groups {
		database.InsertEntity(&model.Entity{
			Id:         uint32(k),
			System:     model.SystemTypeDMR,
			Type:       model.EntityTalkgroup,
			Name:       v,
			Country:    utils.GetCountryNameForID(k),
			CountryISO: utils.GetCountryCodeForID(k),
		})
	}

	users := dmrmarc.FetchUsers()
	for _, u := range users {
		id, _ := strconv.Atoi(u.RadioId)
		database.InsertEntity(&model.Entity{
			Id:         uint32(id),
			System:     model.SystemTypeDMR,
			Type:       model.EntityHam,
			Callsign:   u.Callsign,
			Name:       u.Name,
			Surname:    u.Surname,
			Country:    u.Country,
			CountryISO: utils.GetCountryCodeForID(id),
			City:       u.City,
			State:      u.State,
		})
	}
	log.Print("End!")
}
