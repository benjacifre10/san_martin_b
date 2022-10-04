// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* Correlative model for the mongo DB */
type Correlative struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Year string `bson:"year" json:"year,omitempty"`
	Correlative []string `bson:"correlative" json:"correlative,omitempty"`
	SubjectXStudyPlanId string `bson:"subjectxstudyplanid" json:"subjectXStudyPlanId,omitempty"`
}

