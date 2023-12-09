package main

import (
	"be-project/config"
	handler "be-project/http"
	"be-project/repository"
	"be-project/routes"
	"be-project/service"
)

func main() {
	getDB := config.RunMigration()
	orderRepository := repository.NewOrderRepository(getDB)
	userRepository := repository.NewUserRepository(getDB)
	roleRepository := repository.NewRoleRepository(getDB)
	artikelRepository := repository.NewArtikelRepository(getDB)
	eventRepository := repository.NewEventRepository(getDB)

	userService := service.NewUserService(userRepository)
	orderService := service.NewOrderService(orderRepository)
	roleService := service.NewRoleService(roleRepository)
	artikelService := service.NewArtikelService(artikelRepository)
	eventService := service.NewEventServices(eventRepository)

	userHandler := handler.NewUserHandler(userService)
	orderHandler := handler.NewOrderHandler(orderService)
	roleHandler := handler.NewRoleHandler(roleService)
	artikelHandler := handler.NewArtikelHandler(artikelService)
	eventHandler := handler.NewEventHandler(eventService)

	routes.Run(userHandler, orderHandler, roleHandler, artikelHandler, eventHandler)
}