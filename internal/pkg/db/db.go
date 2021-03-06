package db

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/yeyus/dmr-entities/pkg/api"
)

type DBManager struct {
	host     string
	user     string
	password string
	database string
	handler  *pg.DB
}

func GetDBManager(host string, user string, password string, database string, debug bool) *DBManager {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     host,
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
	for _, model := range []interface{}{&api.Entity{}} {
		log.Printf("[DB] Creating table for model %T", model)
		err := db.handler.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db DBManager) DropSchema() error {
	for _, model := range []interface{}{&api.Entity{}} {
		log.Printf("[DB] Dropping table for model %T", model)
		err := db.handler.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (db DBManager) InsertEntity(entity *api.Entity) error {
	return db.handler.Insert(entity)
}

func (db *DBManager) GetEntity(id uint64) (*api.Entity, error) {
	entity := &api.Entity{
		Id: id,
	}
	err := db.handler.Select(entity)

	return entity, err
}
