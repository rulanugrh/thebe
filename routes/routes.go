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

func Run(user portHandler.UserInterface, order portHandler.OrderInterface, role portHandler.RoleInterface, artikel portHandler.ArtikelInterface, event portHandler.EventInterface) {
	conf := config.GetConfig()
	
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.CommonMiddleware)

	router.HandleFunc("/user/register/", user.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login/", user.Login).Methods("POST", "OPTIONS")
	
	routerGroup := router.PathPrefix("/api/").Subrouter()
	routerGroup.Use(middleware.JWTVerify)
	routerGroup.Use(middleware.CommonMiddleware)
	
	routerGroup.HandleFunc("/user/", user.Update).Methods("PUT")
	routerGroup.HandleFunc("/user/{id}", user.Delete).Methods("DELETE")
	// routing for order
	routerGroup.HandleFunc("/order/", order.Create).Methods("POST")
	routerGroup.HandleFunc("/order/{id}", order.Update).Methods("PUT")
	// routing for role
	routerGroup.HandleFunc("/role/{id}", role.FindByID).Methods("GET")
	routerGroup.HandleFunc("/role/{id}", role.Update).Methods("PUT")
	routerGroup.HandleFunc("/order/{user_id}", order.FindByUserID).Methods("GET")
	// routing for artikel
	routerGroup.HandleFunc("/artikel/", artikel.Create).Methods("POST")
	routerGroup.HandleFunc("/artikel/{id}", artikel.Delete).Methods("DELETE")
	routerGroup.HandleFunc("/artikel/{id}", artikel.FindByID).Methods("GET")
	routerGroup.HandleFunc("/artikel/", artikel.FindAll).Methods("GET")
	// routing for user
	// routing for role
	routerGroup.HandleFunc("/event/", event.Create).Methods("POST")
	routerGroup.HandleFunc("/event/{id}", event.Update).Methods("PUT")
	routerGroup.HandleFunc("/event/{id}", event.FindByID).Methods("GET")
	routerGroup.HandleFunc("/event/{id}/submission", event.SubmissionTask).Methods("POST")
	
	
	host := fmt.Sprintf("%s:%s", conf.App.Host, conf.App.Port)
	server := http.Server{
		Addr: host,
		Handler: router,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Cannot running, because: %s", err.Error())
	}
	log.Printf("Server running at: %v", host)
}
