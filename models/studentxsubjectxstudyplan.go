// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* StudentXSubjectXStudyPlan model for the mongo DB */
type StudentXSubjectXStudyPlan struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	FinalNote string `bson:"finalnote" json:finalNote,omitempty`
	Approved bool `bson:"approved" json:"approved"`
	SubjectStudyPlanId string `bson:"subjectstudyplanid" json:"subjectStudyPlanId,omitempty"`
	StudentId string `bson:"studentid" json:"studentId,omitempty"`
}

/***************************************************************/
/***************************************************************/
/* StudentXSubjectXStudyPlanResponse model for the mongo DB */
type StudentXSubjectXStudyPlanResponse struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	FinalNote string `bson:"finalnote" json:finalNote,omitempty`
	Approved bool `bson:"approved" json:"approved"`
	Subject string `bson:"subject" json:"subject,omitempty"`
	Student string `bson:"student" json:"student,omitempty"`
}
