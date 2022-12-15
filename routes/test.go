// Package routes provides ...
package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* Test Routes */
func TestRoutes(router *mux.Router) (*mux.Router) {
	router.HandleFunc("/test", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetTest))).Methods("GET")	
	router.HandleFunc("/test", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertTest))).Methods("POST")	
	router.HandleFunc("/test", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteTest))).Methods("DELETE")	

	return router
}

