package services

import (
	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetStudentTestService call the db to get the test by student */
func GetStudentTestService(ID string) ([]models.StudentXTestResponse, int, error) {
	// call the db
	result, code, err := db.GetStudentTestsDB()
	if err != nil || code != 200 {
		return result, 400, err
	}

	for i := 0; i < len(result); i++ {
		student, err := db.GetStudentsDB()
		for j := 0; j < len(student); j++ {
			if (student[j].ID.Hex() == result[i].Student) {
				result[i].Student = student[j].Name + " " + student[j].Surname
			}
		}
		subjectxstudyplan, err := db.GetSubjectXStudyPlanDB(result[i].Subject)
		subject, err := db.GetSubjectDB(subjectxstudyplan.SubjectId)
		result[i].Subject = subject.Name
		if err != nil || code != 200 {
			return result, 400, err
		}
	}

	return result, code, nil
}

/***************************************************************/
/***************************************************************/
/* InsertStudentXTestService call the db to insert the test */
func InsertStudentXTestService(s models.StudentXTest) (string, int, error) {
	// check if the test is empty
	if len(s.TestId) == 0 {
		return "No puede inscribirse sin asociar el examen", 199, nil
	}

	// check if the studentsubjectstudyplan is empty
	if len(s.StudentSubjectStudyPlanId) == 0 {
		return "No puede registrar el examen sin el alumno", 199, nil
	}

	row := models.StudentXTest {
		TestId: s.TestId,
		Note: s.Note,
		StudentSubjectStudyPlanId: s.StudentSubjectStudyPlanId,
	}

	msg, err := db.InsertStudentXTestDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateTestNoteService update the note */
func UpdateTestNoteService(s models.StudentXTest) (string, int, error) {

	_, err := db.UpdateTestNoteDB(s)
	if err != nil {
		return "Hubo un error al actualizar la nota del examen en la base", 400, err
	}

	return "La nota del examen se actualizo correctamente", 200, nil
}

