package main

import (
	"be-project/helper"
	handler "be-project/http"
	"be-project/repository"
	"be-project/routes"
	"be-project/service"
)

func main() {
	getDB := helper.RunMigration()
	snaps, env, serverKey := helper.InitMidtrans()
	// conf := config.GetConfig()

	orderRepository := repository.NewOrderRepository(getDB)
	userRepository := repository.NewUserRepository(getDB)
	roleRepository := repository.NewRoleRepository(getDB)
	artikelRepository := repository.NewArtikelRepository(getDB)
	eventRepository := repository.NewEventRepository(getDB)
	paymentRepository := repository.NewPaymentRepository(getDB)

	userService := service.NewUserService(userRepository)
	orderService := service.NewOrderService(orderRepository)
	roleService := service.NewRoleService(roleRepository)
	artikelService := service.NewArtikelService(artikelRepository)
	eventService := service.NewEventServices(eventRepository)
	paymentService := service.NewPaymentService(paymentRepository, env, serverKey, snaps, orderRepository)

	userHandler := handler.NewUserHandler(userService)
	orderHandler := handler.NewOrderHandler(orderService)
	roleHandler := handler.NewRoleHandler(roleService)
	artikelHandler := handler.NewArtikelHandler(artikelService)
	eventHandler := handler.NewEventHandler(eventService)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	routes.Run(userHandler, orderHandler, roleHandler, artikelHandler, eventHandler, paymentHandler)
}
