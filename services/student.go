package services

import (
	"regexp"
	"time"

	"github.com/benjacifre10/san_martin_b/db"
	"github.com/benjacifre10/san_martin_b/models"
)

/***************************************************************/
/***************************************************************/
/* InsertStudentService call the db to insert the student */
func InsertStudentService(s models.Student) (string, int, error) {
	// check if the student is empty
	if len(s.Name) == 0 || len(s.Surname) == 0 || len(s.IdentityNumber) == 0 {
		return "No puede registrar el alumno con nombre o dni vacio", 199, nil
	}

	// check if the degree is empty
	if len(s.DegreeId) == 0 {
		return "No puede registrar plan de estudio sin una carrera asociada", 199, nil
	}

	// verify if the name or surname has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, s.Name + s.Surname)
	if anyNumber == true {
		return "No puede registrar el nombre o apellido del alumno con numeros", 199, errRegexp
	}

	// verify if the identification number has already exists
	_, check, errorCheck := db.CheckExistStudent(s.IdentityNumber)
	if check == true {
		return "Ya existe ese alumno en el sistema", 199, errorCheck
	}

	row := models.Student {
		Name: s.Name,
		Surname: s.Surname,
		IdentityNumber: s.IdentityNumber,
		Address: s.Address,
		Phone: s.Phone,
		Cuil: s.Cuil,
		Arrears: s.Arrears,
		State: s.State,
		UserId: s.UserId,
		DegreeId: s.DegreeId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	msg, err := db.InsertStudentDB(row)
	if err != nil {
		return msg, 400, err
	}

	return msg, 201, nil
}

/***************************************************************/
/***************************************************************/
/* GetStudentsService call the db to get the students */
func GetStudentsService() ([]*models.Student, int, error) {
	// call the db
	result, err := db.GetStudentsDB()
	if err != nil {
		return result, 400, err
	}

	return result, 200, nil
}

/***************************************************************/
/***************************************************************/
/* GetStudentService call the db to get the student by id*/
func GetStudentService(ID string) (models.Student, int, error) {
	// call the db
	result, err := db.GetStudentByIdDB(ID)
	if err != nil {
		return result, 400, err
	}

	return result, 200, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateStudentService update the student */
func UpdateStudentService(s models.Student) (string, int, error) {
	if len(s.Name) == 0 || len(s.Surname) == 0 {
		return "El estudiante debe tener nombre y apellido", 199, nil
	}

	// verify if the type has any number
	anyNumber, errRegexp := regexp.MatchString(`\d+`, s.Name)
	if anyNumber == true {
		return "No puede actualizar el nombre del estudiante con numeros", 199, errRegexp
	}

	// verify if the type has any number
	anyNumber, errRegexp = regexp.MatchString(`\d+`, s.Surname)
	if anyNumber == true {
		return "No puede actualizar el apellido del estudiante con numeros", 199, errRegexp
	}

	_, err := db.UpdateStudentDB(s)
	if err != nil {
		return "Hubo un error al actualizar el estudiante en la base", 400, err
	}

	return "El estudiante se actualizo correctamente", 200, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateStudentStatusService update the student active */
func UpdateStudentStatusService(s models.Student) (string, int, error) {

	_, err := db.UpdateStatusStudentDB(s)
	if err != nil {
		return "Hubo un error al actualizar el estado del estudiante en la base", 400, err
	}

	return "El estado del estudiante se actualizo correctamente", 200, nil
}
