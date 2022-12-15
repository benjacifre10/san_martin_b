// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* Test model for the mongo DB */
type Test struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	TestDate string `bson:"testdate" json:testDate,omitempty`
	Sheet string `bson:"sheet" json:"sheet,omitempty"`
	Form string `bson:"form" json:"form,omitempty"`
	SubjectStudyPlanId string `bson:"subjectstudyplanid" json:"subjectStudyPlanId,omitempty"`
	ProfessorId string `bson:"professorid" json:"professorId,omitempty"`
	TestTypeId string `bson:"testtypeid" json:"testTypeId,omitempty"`
}

/***************************************************************/
/***************************************************************/
/* TestResponse model for the mongo DB */
type TestResponse struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Date string `bson:"date" json:date,omitempty`
	Sheet string `bson:"sheet" json:"sheet,omitempty"`
	Form string `bson:"form" json:"form,omitempty"`
	Subject string `bson:"subject" json:"subject,omitempty"`
	Professor string `bson:"professor" json:"professor,omitempty"`
	Test string `bson:"test" json:"test,omitempty"`
}
