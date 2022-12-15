// Package controllers provides ...
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/benjacifre10/san_martin_b/models"
	"github.com/benjacifre10/san_martin_b/services"
)

/***************************************************************/
/***************************************************************/
/* InsertStudentXSubjectXStudyPlan insert one final note in a subject */
func InsertStudentXSubjectXStudyPlan(w http.ResponseWriter, r *http.Request) {
	var studentxsubjectxstudyplan models.StudentXSubjectXStudyPlan
	err := json.NewDecoder(r.Body).Decode(&studentxsubjectxstudyplan)

	msg, code, err := services.InsertStudentXSubjectXStudyPlanService(studentxsubjectxstudyplan)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al inscribirse en la materia en el plan de estudio. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha inscripto a la materia en el plan de estudio correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* UpdateStudentXSubjectXStudyPlan update the final note in the subject */
func UpdateStudentXSubjectXStudyPlan(w http.ResponseWriter, r *http.Request) {
	var studentxsubjectxstudyplan models.StudentXSubjectXStudyPlan
	err := json.NewDecoder(r.Body).Decode(&studentxsubjectxstudyplan)

	var code int
	var msg string
	msg, code, err = services.UpdateStudentXSubjectXStudyPlanService(studentxsubjectxstudyplan)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar la nota final en la materia. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := models.Response {
		Message: msg,
		Code: code,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* GetStudentXSubjectsXStudyPlan get all the subjects by student */
func GetStudentXSubjectsXStudyPlan(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("studentid")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para mostrar los resultados",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	result, code, err := services.GetStudentXSubjectsXStudyPlanService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al consultar las materias del alumno",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := models.Response {
		Data: result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
