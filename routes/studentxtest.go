// Package routes provides ...
package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* StudentXTest Routes */
func StudentXTestRoutes(router *mux.Router) (*mux.Router) {
	router.HandleFunc("/studentxtest", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetStudentXTest))).Methods("GET")	
	router.HandleFunc("/studentxtest", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertStudentXTest))).Methods("POST")	
	router.HandleFunc("/studentxtest", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.ChangeTestNote))).Methods("PUT")	

	return router
}

