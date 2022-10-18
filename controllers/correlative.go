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
	
	ID := r.URL.Query().Get("studyplanid")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para mostrar los resultados",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	subjects, code2, err2 := services.GetSubjectsXStudyPlanService(ID)
	if err2 != nil || code2 != 200 {
		res := models.Response {
			Message: "Error al consultar las materias por plan de estudio",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	var correlatives []*models.Correlative

	for i := 0; i < len(subjects); i++ {
		corr, code, err := services.GetCorrelativesByStudyPlanService(subjects[i].ID.Hex())
		if err != nil || code != 200 {
			res := models.Response {
				Message: "Error al consultar las correlatividades",
				Code: 400,
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		correlatives = append(correlatives, corr[0])
	}

	res := models.Response {
		Data: correlatives,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* InsertCorrelative insert one subject correlative */
func InsertCorrelative(w http.ResponseWriter, r *http.Request) {
	var correlative models.Correlative
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
/* DeleteCorrelative delete one subject correlative */
func DeleteCorrelative(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("studyplanid")
	if len(ID) < 1 {
		res := models.Response {
			Message: "Falta un parametro para borrar la correlatividad",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	subjects, code2, err2 := services.GetSubjectsXStudyPlanService(ID)
	if err2 != nil || code2 != 200 {
		res := models.Response {
			Message: "Error al consultar las materias por plan de estudio",
			Code: 400,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	var msg string
	var code int	

	for i := 0; i < len(subjects); i++ {
		msg, code, err := services.DeleteCorrelativeService(subjects[i].ID.Hex())
		if err != nil || code != 200 {
			res := models.Response {
				Message: "Error al borrar la correlatividad. " + msg,
				Code: code,
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	res := models.Response {
		Message: msg,
		Code: code,
	}
	json.NewEncoder(w).Encode(res)
}
