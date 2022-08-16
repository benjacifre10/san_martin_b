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
/* GetProfessors get all the professors */
func GetProfessors(w http.ResponseWriter, r *http.Request) {
	result, status := services.GetProfessorsService()
	if status == false {
		res := models.Response {
			Message: "Error al consultar los profesores",
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
/* InsertProfessor insert one professor */
func InsertProfessor(w http.ResponseWriter, r *http.Request) {
	var professor models.Professor
	err := json.NewDecoder(r.Body).Decode(&professor)

	msg, code, err := services.InsertProfessorService(professor)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al insertar el profesor. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha insertado el profesor correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* UpdateProfessor update one professor */
func UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	var professor models.Professor
	err := json.NewDecoder(r.Body).Decode(&professor)

	var code int
	var msg string
	msg, code, err = services.UpdateProfessorService(professor)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar el profesor. " + msg,
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
/* DeleteProfessor delete one professor */
func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para borrar el profesor",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	msg, code, err := services.DeleteProfessorService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al borrar el profesor. " + msg,
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

