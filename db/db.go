package db

import (
	"go-boilerplate/app"
	"go-boilerplate/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"

	// Add because the db connection is separated from main.go
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseConnection struct {
	Test struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}

	Development struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}

	Production struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}
}

// Connect executes all command lines variables that changes database data and returns the connection
func Connect() (db *gorm.DB) {
	var err error

	DBConnection := DatabaseConnection{}
	configor.Load(&DBConnection, app.Config.DatabaseFile)

	switch app.Config.Environment {
	case "test":
		db, err = gorm.Open(DBConnection.Test.Adapter, DBConnection.Test.Database)
	case "development":
		db, err = gorm.Open(DBConnection.Development.Adapter, DBConnection.Development.Database)
	default:
		db, err = gorm.Open(DBConnection.Production.Adapter, DBConnection.Production.Database)
	}

	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(app.Config.ShowSQL)

	if app.Config.Migrate {
		migrate(db)
	}

	return
}

// Instance returns the db singleton connection
func Instance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

// migrate rebuild the database
func migrate(db *gorm.DB) {
	db.DropTableIfExists(
		&models.Artist{},
		&models.User{},
	)

	db.AutoMigrate(
		&models.Artist{},
		&models.User{},
	)
}
