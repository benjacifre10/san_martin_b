package services

import (
	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetTestService call the db to get the test */
func GetTestService() ([]models.TestResponse, int, error) {
	// call the db
	result, code, err := db.GetTestsDB()
	if err != nil || code != 200 {
		return result, 400, err
	}

	for i := 0; i < len(result); i++ {
		subject, err := db.GetSubjectDB(result[i].Subject)
		if err != nil || code != 200 {
			return result, 400, err
		}
		result[i].Subject = subject.Name
	}

	return result, code, nil
}

/***************************************************************/
/***************************************************************/
/* InsertTestService call the db to insert the test */
func InsertTestService(t models.Test) (string, int, error) {
	// check if the testtype is empty
	if len(t.TestTypeId) == 0 {
		return "No puede registrar el examen sin el tipo de examen", 199, nil
	}

	// check if the subjectstudyplan is empty
	if len(t.SubjectStudyPlanId) == 0 {
		return "No puede registrar el examen sin la materia", 199, nil
	}

	// check if the professor is empty
	if len(t.ProfessorId) == 0 {
		return "No puede registrar el examen sin el profesor", 199, nil
	}

	row := models.Test {
		TestDate: t.TestDate,
		Sheet: t.Sheet,
		Form: t.Form,
		SubjectStudyPlanId: t.SubjectStudyPlanId,
		ProfessorId: t.ProfessorId,
		TestTypeId: t.TestTypeId,
	}

	msg, err := db.InsertTestDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteTestService delete the test */
func DeleteTestService(IDTest string) (string, int, error) {
	// find the test type
	_, errFind := db.GetTestDB(IDTest)
	if errFind != nil {
		return "Hubo un error al intentar localizar el examen a borrar en la base", 400, errFind
	}

	err := db.DeleteTestDB(IDTest)
	if err != nil {
		return "Hubo un error al intentar borrar el examen en la base", 400, err
	}

	return "El examen se borro correctamente", 200, nil
}

