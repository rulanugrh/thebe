package config

import (
	"be-project/entity/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
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


	Midtrans struct {
		EnvironmentType string
		Production struct {
			Client string
			Server string
		}
		Sandbox struct {
			Client string
			Server string
		}
		PaymentAppendURL string
		PaymentOverrideURL string
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

func InitMidtrans() ( snap.Client, midtrans.EnvironmentType, string )  {
	var snap snap.Client
	conf := GetConfig()
	if conf.Midtrans.EnvironmentType == "Sandbox" {
		midtrans.Environment = midtrans.Sandbox
		midtrans.ServerKey = conf.Midtrans.Sandbox.Server
		snap.Env = midtrans.Environment
		snap.ServerKey = conf.Midtrans.Sandbox.Server
	} else {
		midtrans.Environment = midtrans.Production
		midtrans.ServerKey = conf.Midtrans.Production.Server
		snap.Env = midtrans.Environment
		snap.ServerKey = conf.Midtrans.Production.Server
	}

	snap.New(midtrans.ServerKey, midtrans.Environment)
	return snap, midtrans.Environment, midtrans.ServerKey

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

	errFind := getDB.Where("name = ?", adminRole.Name).Find(&adminRole).Error
	if errFind == nil {
		log.Printf("Cannot create because role has been created")
	}

	errFind = getDB.Where("name = ?", pesertaRole.Name).Find(&pesertaRole).Error
	if errFind == nil {
		log.Printf("Cannot create because role has been created")
	}
	
	errFind = getDB.Where("f_name = ?", adminUser.FName).Find(&adminUser).Error
	if errFind == nil {
		log.Printf("Cannot create because user has been created")
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
		conf.Midtrans.Sandbox.Client = ""
		conf.Midtrans.Sandbox.Server = ""
		conf.Midtrans.Production.Client = ""
		conf.Midtrans.Production.Server = ""
		conf.Secret = ""
		conf.Midtrans.EnvironmentType = ""
		conf.Midtrans.PaymentAppendURL = ""
		conf.Midtrans.PaymentOverrideURL = ""

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
	conf.Midtrans.Sandbox.Client = os.Getenv("SANDBOX_CLIENT")
	conf.Midtrans.Sandbox.Server = os.Getenv("SANDBOX_SERVER")
	conf.Midtrans.Production.Client = os.Getenv("PRODUCTION_CLIENT")
	conf.Midtrans.Production.Client = os.Getenv("PRODUCTION_SERVER")
	conf.Midtrans.EnvironmentType = os.Getenv("MIDTRANS_ENVIRONTMENT")
	conf.Secret = os.Getenv("APP_SECRET")
	conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")
	conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
	conf.Midtrans.PaymentAppendURL = os.Getenv("MIDTRANS_APPEND_URL")
	conf.Midtrans.PaymentAppendURL = os.Getenv("MIDTRANS_OVERRIDE_URL")


	return &conf

}