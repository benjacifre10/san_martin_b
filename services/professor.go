package services

import (
	"regexp"

	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetProfessorsService call the db to get the professors */
func GetProfessorsService() ([]*models.Professor, bool) {
	// call the db
	result, status := db.GetProfessorsDB()
	if status == false {
		return result, status
	}

	return result, status
}

/***************************************************************/
/***************************************************************/
/* InsertProfessorService call the db to insert the professor */
func InsertProfessorService(p models.Professor) (string, int, error) {
	// check if the professor is empty
	if len(p.Name) == 0 || len(p.Surname) == 0 || len(p.IdentityNumber) == 0 {
		return "No puede registrar el profesor con nombre o dni vacio", 199, nil
	}

	// verify if the name or surname has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, p.Name + p.Surname)
	if anyNumber == true {
		return "No puede registrar el nombre o apellido del profesor con numeros", 199, errRegexp
	}

	// verify if the identification number has already exists
	_, check, errorCheck := db.CheckExistProfessor(p.IdentityNumber)
	if check == true {
		return "Ya existe ese profesor en el sistema", 199, errorCheck
	}

	row := models.Professor {
		Name: p.Name,
		Surname: p.Surname,
		IdentityNumber: p.IdentityNumber,
	}

	msg, err := db.InsertProfessorDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateProfessorService update the professor */
func UpdateProfessorService(p models.Professor) (string, int, error) {
	// check if the professor is empty
	if len(p.Name) == 0 || len(p.Surname) == 0 {
		return "El profesor no puede venir vacio", 199, nil
	}

	// verify if the name or surname has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, p.Name + p.Surname)
	if anyNumber == true {
		return "No puede actualizar el nombre o apellido del profesor con numeros", 199, errRegexp
	}

	row := models.Professor {
		ID: p.ID,
		Name: p.Name,
		Surname: p.Surname,
	}

	_, err := db.UpdateProfessorDB(row)
	if err != nil {
		return "Hubo un error al actualizar el profesor en la base", 400, err
	}

	return "El profesor se actualizo correctamente", 200, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteProfessorService delete the professor */
func DeleteProfessorService(IDProfessor string) (string, int, error) {
	// find the professor
	_, errFind := db.GetProfessorDB(IDProfessor)
	if errFind != nil {
		return "Hubo un error al intentar localizar el profesor a borrar en la base", 400, errFind
	}

	err := db.DeleteProfessorDB(IDProfessor)
	if err != nil {
		return "Hubo un error al intentar borrar el profesor en la base", 400, err
	}

	return "El profesor se borro correctamente", 200, nil
}

