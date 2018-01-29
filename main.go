package main

import (
	"github.com/yeyus/dmr-entities/brandmeister"
	"github.com/yeyus/dmr-entities/db"
	"github.com/yeyus/dmr-entities/model"
	"github.com/yeyus/dmr-entities/utils"
	//	"github.com/yeyus/dmr-entities/dmrmarc"
	"log"
)

const POSTGRES_PASSWORD string = "p0stgr35ql"
const POSTGRES_HOST string = "localhost:5432"
const POSTGRES_USERNAME string = "postgres"
const POSTGRES_DATABASE string = "postgres"

func main() {
	database := db.GetDBManager(POSTGRES_HOST, POSTGRES_USERNAME, POSTGRES_PASSWORD, POSTGRES_DATABASE)
	err := database.DropSchema()
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
	log.Print("\nEnd!\n")
}
