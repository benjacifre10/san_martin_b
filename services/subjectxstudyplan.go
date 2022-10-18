package services

import (
	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* GetSubjectsXStudyPlanService call the db to get the subjects by an study plan*/
func GetSubjectsXStudyPlanService(ID string) ([]models.SubjectXStudyPlanResponse, int, error) {
	// call the db
	result, code, err := db.GetSubjectsXStudyPlanDB(ID)
	if err != nil || code != 200 {
		return result, code, err
	}

	return result, code, nil
}

/***************************************************************/
/***************************************************************/
/* InsertSubjectXStudyPlanService call the db to associate the subject with an study plan */
func InsertSubjectXStudyPlanService(s models.SubjectXStudyPlan) (string, int, error) {
	// check if the subject is empty
	if len(s.SubjectId) == 0 {
		return "No puede asociar con una carrera vacia", 199, nil
	}

	// check if the study plan is empty
	if len(s.StudyPlanId) == 0 {
		return "No puede asociar con un plan de estudio vacio", 199, nil
	}

	row := models.SubjectXStudyPlan {
		SubjectId: s.SubjectId,
		StudyPlanId: s.StudyPlanId,
	}

	msg, err := db.InsertSubjectXStudyPlanDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateSubjectXStudyPlanService update the subject in the study plan */
func UpdateSubjectXStudyPlanService(s models.SubjectXStudyPlan) (string, int, error) {
	if len(s.StudyPlanId) == 0 {
		return "El plan de estudio  no puede venir vacio", 199, nil
	}

	if len(s.SubjectId) == 0 {
		return "La materia no puede venir vacia", 199, nil
	}

	_, err := db.UpdateSubjectXStudyPlanDB(s)
	if err != nil {
		return "Hubo un error al actualizar la carrera en el plan de estudio en la base", 400, err
	}

	return "La carrera se actualizo correctamente en el plan de estudio", 200, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteSubjectXStudyPlanService delete the subject */
func DeleteSubjectXStudyPlanService(IDSubjectXStudyPlan string) (string, int, error) {
	// find the subject x study plan
	_, errFind := db.GetSubjectXStudyPlanDB(IDSubjectXStudyPlan)
	if errFind != nil {
		return "Hubo un error al intentar localizar la materia asociada al plan de estudio a borrar en la base", 400, errFind
	}

	err := db.DeleteSubjectXStudyPlanDB(IDSubjectXStudyPlan)
	if err != nil {
		return "Hubo un error al intentar desasociar la materia con el plan de estudio en la base", 400, err
	}

	return "La materia se desasocio correctamente", 200, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteSubjectXStudyPlanByStudyPlanService delete the subject */
func DeleteSubjectXStudyPlanByStudyPlanService(IDSubjectXStudyPlan string) (string, int, error) {

	err := db.DeleteSubjectXStudyPlanByStudyPlanDB(IDSubjectXStudyPlan)
	if err != nil {
		return "Hubo un error al intentar desasociar las materias con el plan de estudio en la base", 400, err
	}

	return "Las materias se desasociaron correctamente", 200, nil
}
