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
/* GetCorrelativesByStudyPlan get all the correlatives from a study plan */
func GetCorrelativesByStudyPlan(w http.ResponseWriter, r *http.Request) {
	result, status := services.GetCorrelativesByStudyPlanService()
	if status == false {
		res := models.Response {
			Message: "Error al consultar las correlatividades",
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
/* InsertCorrelative insert one subject correlative */
func InsertCorrelative(w http.ResponseWriter, r *http.Request) {
	var correlative models.Degree
	err := json.NewDecoder(r.Body).Decode(&correlative)

	msg, code, err := services.InsertCorrelativeService(correlative)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al insertar la correlatividad. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha insertado la correlatividad correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* UpdateCorrelative update one subject correlative */
func UpdateCorrelative(w http.ResponseWriter, r *http.Request) {
	var correlative models.Degree
	err := json.NewDecoder(r.Body).Decode(&correlative)

	var code int
	var msg string
	msg, code, err = services.UpdateCorrelativeService(correlative)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar la correlatividad. " + msg,
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
/* DeleteCorrelative delete one subject correlative */
func DeleteCorrelative(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para borrar la correlatividad",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	msg, code, err := services.DeleteCorrelativeService(ID)
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al borrar la correlatividad. " + msg,
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
