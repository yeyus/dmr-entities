package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/pariz/gountries"
	"github.com/yeyus/dmr-entities/internal/pkg/db"
	"github.com/yeyus/dmr-entities/internal/pkg/utils"
	"github.com/yeyus/dmr-entities/pkg/api"
	"github.com/yeyus/dmr-entities/pkg/providers/brandmeister"
	"github.com/yeyus/dmr-entities/pkg/providers/dmrmarc"
	"github.com/yeyus/dmr-entities/pkg/providers/radioid"
)

type Configuration struct {
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseName     string `envconfig:"DB_NAME"`
	Debug            bool   `envconfig:"DEBUG"`
}

type updater func(*api.Entity) bool

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

	updater := entityUpdater(database)

	fetchNXDNUsers(updater)
	fetchBrandmeisterGroups(updater)
	fetchDMRUsers(updater)
	log.Print("End!")
}

func entityUpdater(database *db.DBManager) updater {
	return func(fetchedEntity *api.Entity) bool {
		database.InsertEntity(fetchedEntity)
		return true
	}
}

func fetchDMRUsers(updater updater) {
	log.Printf("[syncjob][dmr users] starting job")
	users, err := dmrmarc.FetchUsers()
	if err != nil {
		log.Printf("[syncjob][dmr users] error downloading dmr-marc user database: %s", err)
		return
	}

	log.Printf("[syncjob][dmr users] parsed %d dmr users", len(users))
	var updateCount int
	for _, u := range users {
		id, _ := strconv.Atoi(u.RadioId)
		entity := &api.Entity{
			SystemId:   uint64(id),
			System:     api.SystemType_DMR,
			Type:       api.EntityType_HAM,
			Namespace:  api.Namespace_RADIOID,
			Callsign:   u.Callsign,
			Name:       u.Name,
			Surname:    u.Surname,
			Country:    u.Country,
			CountryIso: utils.GetCountryCodeForID(id),
			City:       u.City,
			State:      u.State,
		}

		if updater(entity) {
			updateCount++
		}
	}

	log.Printf("[syncjob][dmr users] updated %d out of %d dmr users", updateCount, len(users))
}

func fetchBrandmeisterGroups(updater updater) {
	log.Printf("[syncjob][bm talkgroups] starting job")
	groups, err := brandmeister.GetGroups()
	if err != nil {
		log.Printf("[syncjob][bm talkgroups] error downloading brandmeister talkgroups database: %s", err)
		return
	}

	var updateCount int
	for k, v := range groups {
		entity := &api.Entity{
			SystemId:   uint64(k),
			System:     api.SystemType_DMR,
			Type:       api.EntityType_TALKGROUP,
			Namespace:  api.Namespace_BRANDMEISTER,
			Name:       v,
			Country:    utils.GetCountryNameForID(k),
			CountryIso: utils.GetCountryCodeForID(k),
		}

		if updater(entity) {
			updateCount++
		}
	}

	log.Printf("[syncjob][bm talkgroups] updated %d out of %d brandmeister talkgroups", updateCount, len(groups))
}

func fetchNXDNUsers(updater updater) {
	log.Printf("[syncjob][nxdn users] starting job")
	users, err := radioid.FetchUsers()
	if err != nil {
		log.Printf("[syncjob][nxdn users] error downloading brandmeister talkgroups database: %s", err)
		return
	}

	countries := gountries.New()
	var updateCount int
	for _, u := range users {
		entity := &api.Entity{
			SystemId:  uint64(u.SystemId),
			System:    api.SystemType_NXDN,
			Type:      api.EntityType_HAM,
			Namespace: api.Namespace_RADIOID,
			Callsign:  u.Callsign,
			Name:      u.Name,
			Surname:   u.Surname,
			City:      u.City,
			State:     u.State,
			Country:   u.Country,
		}

		country, err := countries.FindCountryByName(u.Country)
		if err == nil {
			entity.CountryIso = strings.ToLower(country.Alpha2)
		}

		if updater(entity) {
			updateCount++
		}
	}

	log.Printf("[syncjob][nxdn users] updated %d out of %d nxdn users", updateCount, len(users))
}
