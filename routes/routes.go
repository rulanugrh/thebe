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
	router := mux.NewRouter().StrictSlash(true)
	router.Use(CommonMiddleware)
	router.HandleFunc("/user/register/", user.Register).Methods("POST")
	router.HandleFunc("/user/login/", user.Login).Methods("POST")
	router.HandleFunc("/role/", role.Create).Methods("POST")

	routerGroup := router.PathPrefix("/api/").Subrouter()
	routerGroup.Use(middleware.CommonMiddleware)
	routerGroup.Use(middleware.JWTVerify)

	// routing for user
	routerGroup.HandleFunc("/user/", user.Update).Methods("PUT")
	routerGroup.HandleFunc("/user/{id}", user.Delete).Methods("DELETE")

	// routing for order
	routerGroup.HandleFunc("/order/", order.Create).Methods("POST")
	routerGroup.HandleFunc("/order/{user_id}", order.FindByUserID).Methods("GET")
	routerGroup.HandleFunc("/order/{id}", order.Update).Methods("PUT")

	// routing for role
	routerGroup.HandleFunc("/role/{id}", role.FindByID).Methods("GET")
	routerGroup.HandleFunc("/role/{id}", role.Update).Methods("PUT")

	// routing for artikel
	routerGroup.HandleFunc("/artikel/", artikel.Create).Methods("POST")
	routerGroup.HandleFunc("/artikel/{id}", artikel.FindByID).Methods("GET")
	routerGroup.HandleFunc("/artikel/", artikel.FindAll).Methods("GET")
	routerGroup.HandleFunc("/artikel/{id}", artikel.Delete).Methods("DELETE")

	// routing for role
	routerGroup.HandleFunc("/event/", event.Create).Methods("POST")
	routerGroup.HandleFunc("/event/{id}", event.FindByID).Methods("GET")
	routerGroup.HandleFunc("/event/{id}", event.Update).Methods("PUT")


	conf := config.GetConfig()
	host := fmt.Sprintf("%s:%s", conf.App.Host, conf.App.Port)

	server := http.Server{
		Addr: host,
		Handler: router,
	}


	log.Printf("Server running at: %v", host)
	server.ListenAndServe()
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "*/*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}