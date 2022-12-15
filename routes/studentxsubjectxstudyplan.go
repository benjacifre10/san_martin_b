package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/benjacifre10/san_martin_b/middlewares"
	"github.com/gorilla/mux"
)

/* StudentXSubjectXStudyPlan Routes */
func StudentXSubjectXStudyPlanRoutes(router *mux.Router) (*mux.Router) {

	router.HandleFunc("/studentxsubjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.GetStudentXSubjectsXStudyPlan))).Methods("GET")
	router.HandleFunc("/studentxsubjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.InsertStudentXSubjectXStudyPlan))).Methods("POST")
	router.HandleFunc("/studentxsubjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.UpdateStudentXSubjectXStudyPlan))).Methods("PUT")
//	router.HandleFunc("/subjectxplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteSubjectXStudyPlan))).Methods("DELETE")
//	router.HandleFunc("/subjectxplan/studyplan", middlewares.DbCheck(middlewares.ValidatedJWT(controllers.DeleteSubjectXStudyPlanByStudyPlan))).Methods("DELETE")

	return router
}


