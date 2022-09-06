package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* SubjectXStudyPlan Routes */
func SubjectXStudyPlanRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetCorrelativesByStudyPlan))).Methods("GET")
	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertCorrelative))).Methods("POST")
	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.UpdateCorrelative))).Methods("PUT")
	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteCorrelative))).Methods("DELETE")

	return router
}

