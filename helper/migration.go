package helper

import (
	"be-project/config"
	"be-project/entity/domain"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunMigration() *gorm.DB {
	conf := config.GetConfig()

	getDB := config.GetConnection()
	errs := getDB.AutoMigrate(&domain.Order{}, &domain.Roles{}, &domain.User{}, &domain.Event{}, &domain.Artikel{}, &domain.Submission{}, &domain.Payment{}, &domain.Transaction{})
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

	bytes, err := bcrypt.GenerateFromPassword([]byte(conf.Admin.Password), 14)
	if err != nil {
		log.Printf("Cant generate hash password: %v", err)
	}

	adminUser := domain.User{
		Name:      "Admin",
		Telephone: "_",
		Address:   "-",
		Email:     conf.Admin.Email,
		Password:  string(bytes),
		RoleID:    1,
	}

	errFind := getDB.Where("name = ?", adminRole.Name).Find(&adminRole).Error
	if errFind != nil {
		log.Printf("Cannot create because role has been created")
	}

	errFind = getDB.Where("name = ?", pesertaRole.Name).Find(&pesertaRole).Error
	if errFind != nil {
		log.Printf("Cannot create because role has been created")
	}

	errFind = getDB.Where("name = ?", adminUser.Name).Find(&adminUser).Error
	if errFind != nil {
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