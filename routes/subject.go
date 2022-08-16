package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* Subject Routes */
func SubjectRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/subject", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetSubjects))).Methods("GET")
	router.HandleFunc("/subject", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertSubject))).Methods("POST")
	router.HandleFunc("/subject", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.UpdateSubject))).Methods("PUT")
	router.HandleFunc("/subject", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteSubject))).Methods("DELETE")

	return router
}

