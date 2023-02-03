package routes

import (
	"github.com/Funmi4194/myMod/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(route *mux.Router) {

	route.HandleFunc("/count/", controllers.CreateDocument()).Methods("POST")

}
