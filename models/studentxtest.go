// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* StudentXTest model for the mongo DB */
type StudentXTest struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	TestId string `bson:"testid" json:testId,omitempty`
	Note string `bson:"note" json:"note"`
	StudentSubjectStudyPlanId string `bson:"studentsubjectstudyplanid" json:"studentSubjectStudyPlanId,omitempty"`
}

/***************************************************************/
/***************************************************************/
/* StudentXTestResponse model for the mongo DB */
type StudentXTestResponse struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Test string `bson:"test" json:"test"`
	Note string `bson:"note" json:"note"`
	Student string `bson:"student" json:"student"`
	Subject string `bson:"subject" json:"subject,omitempty"`
	Date string `bson:"date" json:"date"`
}
