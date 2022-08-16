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
/* GetSubjects get all the academy subjects */
func GetSubjects(w http.ResponseWriter, r *http.Request) {
	result, status := services.GetSubjectsService()
	if status == false {
		res := models.Response {
			Message: "Error al consultar las materias",
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
/* InsertSubject insert one academy subject */
func InsertSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)

	msg, code, err := services.InsertSubjectService(subject)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al insertar la materia. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha insertado la materia correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* UpdateSubject update one academy subject */
func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)

	var code int
	var msg string
	msg, code, err = services.UpdateSubjectService(subject)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar la materia. " + msg,
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
/* DeleteSubject delete one academy subject */
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para borrar la materia",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	msg, code, err := services.DeleteSubjectService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al borrar la materia. " + msg,
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
