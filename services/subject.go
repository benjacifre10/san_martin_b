package services

import (
	"regexp"

	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetSubjectsService call the db to get the subjects */
func GetSubjectsService() ([]*models.Subject, bool) {
	// call the db
	result, status := db.GetSubjectsDB()
	if status == false {
		return result, status
	}

	return result, status
}

/***************************************************************/
/***************************************************************/
/* InsertSubjectService call the db to insert the subject */
func InsertSubjectService(s models.Subject) (string, int, error) {
	// check if the subject is empty
	if len(s.Name) == 0 {
		return "No puede registrar la materia vacia", 199, nil
	}

	// verify if the name has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, s.Name)
	if anyNumber == true {
		return "No puede registrar la materia con numeros", 199, errRegexp
	}

	// verify if the name has already exists
	_, check, errorCheck := db.CheckExistSubject(s.Name)
	if check == true {
		return "Ya existe esa materia en el sistema", 199, errorCheck
	}

	row := models.Subject {
		Name: s.Name,
		ProfessorId: s.ProfessorId,
		ShiftId: s.ShiftId,
		PursueTypeId: s.PursueTypeId,
		CreditHours: s.CreditHours,
		Days: s.Days,
		From: s.From,
		To: s.To,
	}

	msg, err := db.InsertSubjectDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateSubjectService update the academy subject */
func UpdateSubjectService(s models.Subject) (string, int, error) {
	if len(s.Name) == 0 {
		return "La materia no puede venir vacia", 199, nil
	}

	// verify if the name has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, s.Name)
	if anyNumber == true {
		return "No puede actualizar la materia con numeros", 199, errRegexp
	}

	// verify if the name has already exists
	_, check, errorCheck := db.CheckExistSubject(s.Name)
	if check == true {
		return "Ya existe esa materia en el sistema", 199, errorCheck
	}

	_, err := db.UpdateSubjectDB(s)
	if err != nil {
		return "Hubo un error al actualizar la materia en la base", 400, err
	}

	return "La materia se actualizo correctamente", 200, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteSubjectService delete the academy subject */
func DeleteSubjectService(IDSubject string) (string, int, error) {
	// find the subject
	_, errFind := db.GetSubjectDB(IDSubject)
	if errFind != nil {
		return "Hubo un error al intentar localizar la materia a borrar en la base", 400, errFind
	}

	err := db.DeleteSubjectDB(IDSubject)
	if err != nil {
		return "Hubo un error al intentar borrar la materia en la base", 400, err
	}

	return "La materia se borro correctamente", 200, nil
}
