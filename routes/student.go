package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* Student Routes */
func StudentRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/student", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetStudent))).Methods("GET")
	router.HandleFunc("/student", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertStudent))).Methods("POST")
	router.HandleFunc("/student", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.UpdateStudent))).Methods("PUT")
	router.HandleFunc("/student/state", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.ChangeActiveStudent))).Methods("PUT")

	return router
}
