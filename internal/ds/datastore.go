package ds

import (
	"github.com/eburyagin/bizone-acc/internal/cfg"
	"github.com/go-pg/pg"
)

func Connect(config *cfg.Config) (*pg.DB, error) {
	ds := pg.Connect(&pg.Options{
		Addr:     config.Datastore.Addr,
		User:     config.Datastore.User,
		Password: config.Datastore.Password,
		Database: config.Datastore.Database,
	})
	var n int
	_, err := ds.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil || n != 1 {
		return ds, err
	}
	return ds, nil
}

func Disconnect(ds *pg.DB) error {
	ds.Close()
	return nil
}
