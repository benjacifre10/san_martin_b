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
/* GetTest get all the test */
func GetTest(w http.ResponseWriter, r *http.Request) {
	result, code, err := services.GetTestService()
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al consultar los examenes",
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
/* InsertTest insert one test */
func InsertTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	err := json.NewDecoder(r.Body).Decode(&test)

	msg, code, err := services.InsertTestService(test)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al insertar el examen. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha insertado el examen correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* DeleteTest delete one test */
func DeleteTest(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para borrar el examen",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	msg, code, err := services.DeleteTestService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al borrar el examen. " + msg,
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
