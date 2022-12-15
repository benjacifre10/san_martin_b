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

