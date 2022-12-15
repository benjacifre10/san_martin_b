package services

import (
	"strconv"

	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* InsertStudentXSubjectXStudyPlanService call the db to associate the final note of subject with an study plan */
func InsertStudentXSubjectXStudyPlanService(s models.StudentXSubjectXStudyPlan) (string, int, error) {
	// check if the subject is empty
	if len(s.SubjectStudyPlanId) == 0 {
		return "No puede asociar con una materia vacia", 199, nil
	}

	// check if the study plan is empty
	if len(s.StudentId) == 0 {
		return "No puede asociar con un estudiante nulo", 199, nil
	}

	row := models.StudentXSubjectXStudyPlan {
		SubjectStudyPlanId: s.SubjectStudyPlanId,
		StudentId: s.StudentId,
		FinalNote: s.FinalNote,
		Approved: s.Approved,
	}

	msg, err := db.InsertStudentXSubjectXStudyPlanDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateStudentXSubjectXStudyPlanService update the final note in the subject */
func UpdateStudentXSubjectXStudyPlanService(s models.StudentXSubjectXStudyPlan) (string, int, error) {
	if len(s.ID) == 0 {
		return "La materia no puede venir vacia", 199, nil
	}

	note, errNote := strconv.ParseInt(s.FinalNote, 10, 0)
	if errNote != nil {
		return "Hubo un error en la nota final", 400, errNote
	}

	if (note >= 4) {
		s.Approved = true
	} else {
		s.Approved = false
	}

	_, err := db.UpdateStudentXSubjectXStudyPlanDB(s)
	if err != nil {
		return "Hubo un error al actualizar la nota final de la materia en la base", 400, err
	}

	return "La nota final de la materia se actualizo correctamente", 200, nil
}

/***************************************************************/
/***************************************************************/
/* GetStudentXSubjectsXStudyPlanService call the db to get the subjects by an student*/
func GetStudentXSubjectsXStudyPlanService(ID string) ([]models.StudentXSubjectXStudyPlanResponse, int, error) {
	// call the db
	result, code, err := db.GetStudentXSubjectsXStudyPlanDB(ID)
	if err != nil || code != 200 {
		return result, code, err
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
