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
/* GetSubjects get all the subjects by study plan */
func GetSubjectsXStudyPlan(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("studyplanid")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para mostrar los resultados",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	result, code, err := services.GetSubjectsXStudyPlanService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al consultar las materias del plan de estudio",
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

/***************************************************************/
/***************************************************************/
/* InsertSubjectXStudyPlan insert one subject in a study plan */
func InsertSubjectXStudyPlan(w http.ResponseWriter, r *http.Request) {
	var subjectxstudyplan models.SubjectXStudyPlan
	err := json.NewDecoder(r.Body).Decode(&subjectxstudyplan)

	msg, code, err := services.InsertSubjectXStudyPlanService(subjectxstudyplan)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al asociar la materia en el plan de estudio. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha asociado la materia en el plan de estudio correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* UpdateSubjectXStudyPlan update the subject in the study plan */
func UpdateSubjectXStudyPlan(w http.ResponseWriter, r *http.Request) {
	var subjectxstudyplan models.SubjectXStudyPlan
	err := json.NewDecoder(r.Body).Decode(&subjectxstudyplan)

	var code int
	var msg string
	msg, code, err = services.UpdateSubjectXStudyPlanService(subjectxstudyplan)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar la carrera en el plan de estudio. " + msg,
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
/* DeleteSubjectXStudyPlan delete one subject per study plan */
func DeleteSubjectXStudyPlan(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para desasociar la materia con el plan de estudio",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	msg, code, err := services.DeleteSubjectXStudyPlanService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al desasociar la materia con el plan de estudio. " + msg,
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
