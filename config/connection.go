package config

import (
	"be-project/entity/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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
		AllowOrigin string
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

	Admin struct {
		Password string
		Email    string
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

	DB = db
	log.Print("Success connect to database")
	return db
}

func RunMigration() *gorm.DB {
	config := GetConfig()

	getDB := GetConnection()
	errs := getDB.AutoMigrate(&domain.Order{}, &domain.Roles{}, &domain.User{}, &domain.Event{}, &domain.Artikel{}, &domain.DelegasiParticipant{}, &domain.SubmissionTask{}, &domain.Payment{}, &domain.Transaction{})
	if errs != nil {
		log.Printf("Cannot migration, because: %s", errs.Error())
	}

	adminRole := domain.Roles{
		Name:        "administrator",
		Description: "ini adalah role admin",
	}

	pesertaRole := domain.Roles{
		Name:        "peserta",
		Description: "ini adalah role peserta",
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(config.Admin.Password), 14)
	if err != nil {
		log.Printf("Cant generate hash password: %v", err)
	}


	adminUser := domain.User{
		FName:     "Admin",
		LName:     "IAI",
		Telephone: "_",
		Address:   "-",
		Email:     config.Admin.Email,
		Password:  string(bytes),
		RoleID:    1,
	}

	errAdminRole := getDB.Create(&adminRole).Error
	if errAdminRole != nil {
		log.Printf("Cannot create role admin: %s", errAdminRole.Error())
	}

	errpesertaRoles := getDB.Create(&pesertaRole).Error
	if errpesertaRoles != nil {
		log.Printf("Cannot create role peserta: %s", errpesertaRoles.Error())
	}

	errAdmin := getDB.Create(&adminUser).Error
	if errAdmin != nil {
		log.Printf("Cnnot create administrator: %s", errAdmin.Error())
	}

	log.Println("Success migration and create roles & administrator")
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
	conf.App.AllowOrigin = os.Getenv("APP_ORIGIN")

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
	conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")
	conf.Admin.Email = os.Getenv("ADMIN_EMAIL")

	return &conf

}
