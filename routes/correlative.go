package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* Correlative Routes */
func CorrelativeRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/correlative", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetCorrelativesByStudyPlan))).Methods("GET")
	router.HandleFunc("/correlative", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertCorrelative))).Methods("POST")
	router.HandleFunc("/correlative", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteCorrelative))).Methods("DELETE")

	return router
}
