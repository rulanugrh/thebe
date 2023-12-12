package config

import (
	"be-project/entity/domain"
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

	Secret string
}

var app *Config

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
	return db
}

func RunMigration() *gorm.DB {
	getDB := GetConnection()
	getDB.AutoMigrate(&domain.Order{}, &domain.Roles{}, &domain.User{}, &domain.Event{}, &domain.Artikel{}, &domain.DelegasiParticipant{}, &domain.SubmissionTask{}, &domain.Payment{}, &domain.Transaction{})

	adminRole := domain.Roles{
		Name:        "administrator",
		Description: "ini adalah role admin",
	}

	pesertaRole := domain.Roles{
		Name:        "peserta",
		Description: "ini adalah role peserta",
	}

	adminUser := domain.User{
		FName:     "Admin",
		LName:     "IAI",
		Telephone: "_",
		Address:   "-",
		Email:     "admin@admin.co.id",
		Password:  "admin123!!",
		RoleID:    2,
	}

	getDB.Create(&adminRole)
	getDB.Create(&pesertaRole)
	getDB.Create(&adminUser)

	return getDB
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
		conf.Secret = ""

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
	conf.Secret = os.Getenv("APP_SECRET")

	return &conf

}
