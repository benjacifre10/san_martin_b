package services

import (
	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetCorrelativesByStudyPlanService call the db to get the correlatives */
func GetCorrelativesByStudyPlanService(id string) ([]*models.Correlative, int, error) {
	// call the db
	result, code, err := db.GetCorrelativesByStudyPlanDB(id)
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
/* DeleteCorrelativeService delete the subject correlative */
func DeleteCorrelativeService(idSubjectXStudyPlan string) (string, int, error) {
	err := db.DeleteCorrelativeDB(idSubjectXStudyPlan)
	if err != nil {
		return "Hubo un error al actualizar las correlatividades en la base", 400, err
	}

	return "Las correlatividad se actualizo correctamente", 200, nil
}
