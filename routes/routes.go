package routes

import (
	"be-project/config"
	portHandler "be-project/http/port"
	"be-project/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(user portHandler.UserInterface, order portHandler.OrderInterface, role portHandler.RoleInterface, artikel portHandler.ArtikelInterface) error {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.CommonMiddleware)
	router.HandleFunc("/register/", user.Register)
	router.HandleFunc("/login/", user.Login)

	routerGroup := router.PathPrefix("/api/").Subrouter()
	routerGroup.Use(middleware.CommonMiddleware)

	// routing for user
	routerGroup.HandleFunc("/user/", user.Update)
	routerGroup.HandleFunc("/user/{id}", user.Delete)

	// routing for order
	routerGroup.HandleFunc("/order/", order.Create)
	routerGroup.HandleFunc("/order/{user_id}", order.FindByUserID)
	routerGroup.HandleFunc("/order/{id}", order.Update)

	// routing for role
	routerGroup.HandleFunc("/role/", role.Create)
	routerGroup.HandleFunc("/role/{id}", role.FindByID)
	routerGroup.HandleFunc("/role/{id}", role.Update)

	// routing for artikel
	routerGroup.HandleFunc("/role/", artikel.Create)
	routerGroup.HandleFunc("/role/{id}", artikel.FindByID)
	routerGroup.HandleFunc("/role/", artikel.FindAll)
	routerGroup.HandleFunc("/role/{id}", artikel.Delete)


	conf := config.GetConfig()
	host := fmt.Sprintf("%s:%s", conf.App.Host, conf.App.Port)

	server := http.Server{
		Addr: host,
		Handler: router,
	}


	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Cannot running http server, because: %s", err.Error())
		return err
	}

	log.Printf("Server running at: %v", host)
	return nil

}