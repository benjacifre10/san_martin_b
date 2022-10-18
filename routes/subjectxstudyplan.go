package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* SubjectXStudyPlan Routes */
func SubjectXStudyPlanRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetSubjectsXStudyPlan))).Methods("GET")
	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertSubjectXStudyPlan))).Methods("POST")
	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.UpdateSubjectXStudyPlan))).Methods("PUT")
	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteSubjectXStudyPlan))).Methods("DELETE")
	router.HandleFunc("/subjectxplan/studyplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteSubjectXStudyPlanByStudyPlan))).Methods("DELETE")

	return router
}

