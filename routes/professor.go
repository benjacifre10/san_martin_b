package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* Professor Routes */
func ProfessorRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/professor", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetProfessors))).Methods("GET")
	router.HandleFunc("/professor", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertProfessor))).Methods("POST")
	router.HandleFunc("/professor", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.UpdateProfessor))).Methods("PUT")
	router.HandleFunc("/professor", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteProfessor))).Methods("DELETE")

	return router
}
