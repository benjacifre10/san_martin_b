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
/* InsertStudent insert one student */
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)

	msg, code, err := services.InsertStudentService(student)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al insertar el alumno. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha insertado el alumno correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* GetStudent get all the students or one by id */
func GetStudent(w http.ResponseWriter, r *http.Request) {
	Email := r.URL.Query().Get("useremail")
	var res models.Response
	if len(Email) < 1 {
		result, code, err := services.GetStudentsService()
		if err != nil || code != 200 {
			res = models.Response {
				Message: "Error al consultar los estudiantes",
				Code: 400,
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		res = models.Response {
			Code: 200,
			Data: result,
		}
	} else {
		result, code, err := services.GetStudentService(Email)
		if err != nil || code != 200 {
			res = models.Response {
				Message: "Error al consultar el estudiante",
				Code: 400,
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		res = models.Response {
			Code: 200,
			Data: result,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* UpdateStudent update one student */
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)

	var code int
	var msg string
	msg, code, err = services.UpdateStudentService(student)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar el estudiante. " + msg,
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
/* ChangeActiveStudent update status student */
func ChangeActiveStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)

	var code int
	var msg string
	msg, code, err = services.UpdateStudentStatusService(student)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar el estado del estudiante. " + msg,
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
