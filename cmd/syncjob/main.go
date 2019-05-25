package main

import (
	"log"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"github.com/yeyus/dmr-entities/internal/pkg/db"
	"github.com/yeyus/dmr-entities/internal/pkg/utils"
	"github.com/yeyus/dmr-entities/pkg/api"
	"github.com/yeyus/dmr-entities/pkg/providers/brandmeister"
	"github.com/yeyus/dmr-entities/pkg/providers/dmrmarc"
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
		database.InsertEntity(&api.Entity{
			Id:         uint32(k),
			System:     api.SystemType_DMR,
			Type:       api.EntityType_TALKGROUP,
			Name:       v,
			Country:    utils.GetCountryNameForID(k),
			CountryIso: utils.GetCountryCodeForID(k),
		})
	}

	users := dmrmarc.FetchUsers()
	for _, u := range users {
		id, _ := strconv.Atoi(u.RadioId)
		database.InsertEntity(&api.Entity{
			Id:         uint32(id),
			System:     api.SystemType_DMR,
			Type:       api.EntityType_HAM,
			Callsign:   u.Callsign,
			Name:       u.Name,
			Surname:    u.Surname,
			Country:    u.Country,
			CountryIso: utils.GetCountryCodeForID(id),
			City:       u.City,
			State:      u.State,
		})
	}
	log.Print("End!")
}
