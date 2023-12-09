package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}
	App struct {
		Host string
		Port string
	}

	Sandbox struct {
		Client string
		Server string
	}

	Production struct {
		Client string
		Server string
	}
}
var app *Config
var DB *gorm.DB
func GetConnection() *gorm.DB {
	conf := GetConfig()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Shanghai", 
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Cant connect to Postgress because: %s", err.Error())
		return nil
	}

	log.Print("Success connect to database")
	DB = db
	return db
}

func GetConfig() *Config {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *Config {
	conf := Config{}
	if err := godotenv.Load(); err != nil {
		conf.App.Host = "localhost"
		conf.App.Port = "30000"
		conf.Database.Host = "localhost"
		conf.Database.Port = ""
		conf.Database.Name = ""
		conf.Database.User = ""
		conf.Database.Pass = ""
		conf.Sandbox.Client = ""
		conf.Sandbox.Server = ""
		conf.Production.Client = ""
		conf.Production.Server = ""

		return &conf
	}

	conf.App.Host = os.Getenv("APP_HOST")
	conf.App.Port = os.Getenv("APP_PORT")
	conf.Database.Host = os.Getenv("DATABASE_HOST")
	conf.Database.Port = os.Getenv("DATABASE_PORT")
	conf.Database.Name = os.Getenv("DATABASE_NAME")
	conf.Database.Pass = os.Getenv("DATABASE_PASS")
	conf.Database.User = os.Getenv("DATABASE_USER")
	conf.Sandbox.Client = os.Getenv("SANDBOX_CLIENT")
	conf.Sandbox.Server = os.Getenv("SANDBOX_SERVER")
	conf.Production.Client = os.Getenv("PRODUCTION_CLIENT")
	conf.Production.Client = os.Getenv("PRODUCTION_SERVER")

	return &conf

}