package routes

import (
	"github.com/benjacifre10/san_martin_b/controllers"
	"github.com/gorilla/mux"
)

/* HandlerRoutes create the differents routes for my API*/
func HandlerRoutes() (*mux.Router) {
	// create the router
	router := mux.NewRouter()

	router.HandleFunc("/health", controllers.Health).Methods("GET")
	router = CorrelativeRoutes(router)
	router = DegreeRoutes(router)
	router = ProfessorRoutes(router)
	router = PursueTypeRoutes(router)
	router = RoleRoutes(router)
	router = ShiftRoutes(router)
	router = StudentRoutes(router)
	router = StudentXSubjectXStudyPlanRoutes(router)
	router = StudyPlanRoutes(router)
	router = SubjectRoutes(router)
	router = SubjectXStudyPlanRoutes(router)
	router = TestRoutes(router)
	router = TestTypeRoutes(router)
	router = UserRoutes(router)

	return router
}
