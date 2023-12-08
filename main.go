package main

import (
	"be-project/config"
	"be-project/entity/domain"
	handler "be-project/http"
	"be-project/repository"
	"be-project/routes"
	"be-project/service"
	"log"
)

func main() {
	getDB := config.GetConnection()
	getDB.AutoMigrate(&domain.Order{}, &domain.Roles{}, &domain.User{})

	orderRepository := repository.NewOrderRepository(getDB)
	userRepository := repository.NewUserRepository(getDB)
	roleRepository := repository.NewRoleRepository(getDB)
	artikelRepository := repository.NewArtikelRepository(getDB)

	userService := service.NewUserService(userRepository)
	orderService := service.NewOrderService(orderRepository)
	roleService := service.NewRoleService(roleRepository)
	artikelService := service.NewArtikelService(artikelRepository)

	userHandler := handler.NewUserHandler(userService)
	orderHandler := handler.NewOrderHandler(orderService)
	roleHandler := handler.NewRoleHandler(roleService)
	artikelHandler := handler.NewArtikelHandler(artikelService)

	err := routes.Run(userHandler, orderHandler, roleHandler, artikelHandler)
	if err != nil {
		log.Printf("Cannot running , because: %s", err.Error())
	}
}