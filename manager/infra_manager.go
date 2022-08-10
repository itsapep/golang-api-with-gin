package manager

import (
	"fmt"
	"log"

	"github.com/itsapep/golang-api-with-gin/config"
	"github.com/itsapep/golang-api-with-gin/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// infra plays the role of slice in the last case
type Infra interface {
	SqlDb() *gorm.DB
	FilePath() string
}

type infra struct {
	db  *gorm.DB
	cfg config.Config
}

// FilePath implements Infra
func (i *infra) FilePath() string {
	fmt.Println("filepath:", i.cfg.FilePath)
	return i.cfg.FilePath
}

// SqlDb implements Infra
func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName, config.Env)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{
		db:  resource,
		cfg: *config,
	}
}

func initDbResource(dataSourceName string, env string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	defer func() {
		err := recover()
		if err != nil {
			log.Println("application failed to run", err)
		}
	}()

	if env == "migration" {
		migration.DbMigration(db)
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
