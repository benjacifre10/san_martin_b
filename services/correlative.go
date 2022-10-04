package services

import (
	"regexp"

	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetCorrelativesByStudyPlanService call the db to get the correlatives */
func GetCorrelativesByStudyPlanService() ([]*models.Correlative, int, error) {
	// call the db
	result, code, err := db.GetCorrelativesByStudyPlanDB()
	if err != nil || code != 200 {
		return result, code, err
	}

	return result, code, nil
}

/***************************************************************/
/***************************************************************/
/* InsertCorrelativeService call the db to insert the correlative */
func InsertCorrelativeService(c models.Correlative) (string, int, error) {
	// check if the year is empty
	if len(c.Year) == 0 {
		return "No puede registrar la correlatividad sin el anio", 199, nil
	}

	// check if the correlative is empty
	if len(c.Correlative) == 0 {
		return "No puede registrar la correlatividad vacia", 199, nil
	}

	// check if the subject x study plan is empty
	if len(c.SubjectXStudyPlanId) == 0 {
		return "No puede registrar la correlatividad con el plan de estudio vacio", 199, nil
	}

	row := models.Correlative {
		Year: c.Year,
		Correlative: c.Correlative,
		SubjectXStudyPlanId: c.SubjectXStudyPlanId,
	}

	msg, err := db.InsertCorrelativeDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateCorrelativeService update the subject correlative */
func UpdateCorrelativeService(d models.Degree) (string, int, error) {
	if len(d.Name) == 0 {
		return "La carrera no puede venir vacia", 199, nil
	}

	// verify if the type has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, d.Name)
	if anyNumber == true {
		return "No puede actualizar la carrera con numeros", 199, errRegexp
	}

	// verify if the type has already exists
	_, check, errorCheck := db.CheckExistDegree(d.Name)
	if check == true {
		return "Ya existe esa carrera en el sistema", 199, errorCheck
	}

	_, err := db.UpdateDegreeDB(d)
	if err != nil {
		return "Hubo un error al actualizar la carrera en la base", 400, err
	}

	return "La carrera se actualizo correctamente", 200, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteCorrelativeService delete the subject correlative */
func DeleteCorrelativeService(idCorrelative string) (string, int, error) {

	err := db.DeleteProfessorDB(idCorrelative)
	if err != nil {
		return "Hubo un error al actualizar el estado de la carrera en la base", 400, err
	}

	return "El estado de la carrera se actualizo correctamente", 200, nil
}
