package db

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/yeyus/dmr-entities/model"
	"log"
	"time"
)

type DBManager struct {
	host     string
	user     string
	password string
	database string
	handler  *pg.DB
}

func GetDBManager(host string, user string, password string, database string) *DBManager {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     host,
	})

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return &DBManager{
		host:     host,
		user:     user,
		password: password,
		database: database,
		handler:  db,
	}
}

func (db DBManager) CreateSchema() error {
	for _, model := range []interface{}{&model.Entity{}} {
		log.Printf("[DB] Creating table for model %T", model)
		err := db.handler.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db DBManager) DropSchema() error {
	for _, model := range []interface{}{&model.Entity{}} {
		log.Printf("[DB] Creating table for model %T", model)
		err := db.handler.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (db DBManager) InsertEntity(entity *model.Entity) error {
	return db.handler.Insert(entity)
}
