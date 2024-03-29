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

func Run(user portHandler.UserInterface, order portHandler.OrderInterface, role portHandler.RoleInterface, artikel portHandler.ArtikelInterface, event portHandler.EventInterface, payment portHandler.PaymentInterface) {
	conf := config.GetConfig()
	
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.CommonMiddleware)

	router.HandleFunc("/user/register/", user.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login/", user.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/refresh-token/", user.RefreshToken).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/validate-token/", user.ValidateToken).Methods("POST", "OPTIONS")
	router.HandleFunc("/midtrans/payment-callback/", payment.PaymentNotification).Methods("POST")
	
	routerGroup := router.PathPrefix("/api/").Subrouter()
	routerGroup.Use(middleware.JWTVerify)
	
	routerGroup.HandleFunc("/user/", user.Update).Methods("PUT")
	routerGroup.HandleFunc("/user/{id}", user.Delete).Methods("DELETE")
	routerGroup.HandleFunc("/user/logout/", user.Logout).Methods("DELETE")


	routerGroup.HandleFunc("/order/", order.Create).Methods("POST")
	routerGroup.HandleFunc("/order/{id}", order.Update).Methods("PUT")
	
	routerGroup.HandleFunc("/role/{id}", role.FindByID).Methods("GET")
	routerGroup.HandleFunc("/role/{id}", role.Update).Methods("PUT")
	routerGroup.HandleFunc("/order/{user_id}/{uuid}", order.FindByUserIDDetail).Methods("GET")
	routerGroup.HandleFunc("/order/{user_id}", order.FindByUserID).Methods("GET")
	routerGroup.HandleFunc("/order/{uuid}", order.FindByUUID).Methods("GET")

	routerGroup.HandleFunc("/artikel/", artikel.Create).Methods("POST")
	routerGroup.HandleFunc("/artikel/{id}", artikel.Delete).Methods("DELETE")
	routerGroup.HandleFunc("/artikel/{id}", artikel.FindByID).Methods("GET")
	routerGroup.HandleFunc("/artikel/", artikel.FindAll).Methods("GET")

	routerGroup.HandleFunc("/event/", event.Create).Methods("POST")
	routerGroup.HandleFunc("/event/{id}", event.Update).Methods("PUT")
	routerGroup.HandleFunc("/event/{id}", event.FindByID).Methods("GET")
	routerGroup.HandleFunc("/event/upload/submission/", event.SubmissionTask).Methods("POST")
	
	routerGroup.HandleFunc("/order/checkout/", payment.Create).Methods("POST")
	routerGroup.HandleFunc("/order/check/{id}/status", payment.HandlingStatus).Methods("GET")
	
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
