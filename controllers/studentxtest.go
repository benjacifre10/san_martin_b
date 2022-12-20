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
/* GetStudentXTest get all the test */
func GetStudentXTest(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("studentid")
	if len(ID) < 1 {
		result, code, err := services.GetStudentTestService(ID)
//		result, code, err := services.GetAllStudentTestService()
		if err != nil || code != 200 {
			res := models.Response {
				Message: "Error al consultar todos los examenes",
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
	} else {
		result, code, err := services.GetStudentTestService(ID)
		if err != nil || code != 200 {
			res := models.Response {
				Message: "Error al consultar los examenes del alumno",
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

}

/***************************************************************/
/***************************************************************/
/* InsertStudentXTest insert one test */
func InsertStudentXTest(w http.ResponseWriter, r *http.Request) {
	var studentxtest models.StudentXTest
	err := json.NewDecoder(r.Body).Decode(&studentxtest)

	msg, code, err := services.InsertStudentXTestService(studentxtest)
	if err != nil || code != 201 {
		res := models.Response {
			Message: "Error al inscribirse. " + msg,
			Code: code,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := models.Response {
		Message: "Se ha inscripto correctamente",
		Code: code,
		Data: msg,
	}
	json.NewEncoder(w).Encode(res)
}

/***************************************************************/
/***************************************************************/
/* ChangeTestNote update note test */
func ChangeTestNote(w http.ResponseWriter, r *http.Request) {
	var studentxtest models.StudentXTest
	err := json.NewDecoder(r.Body).Decode(&studentxtest)

	var code int
	var msg string
	
	msg, code, err = services.UpdateTestNoteService(studentxtest)
	
	if err != nil || code != 200 {
		res := models.Response {
			Message: "Error al actualizar la nota del examen. " + msg,
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

