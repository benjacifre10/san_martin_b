// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* SubjectXStudyPlan model for the mongo DB */
type SubjectXStudyPlan struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	SubjectId string `bson:"subjectid" json:"subjectId,omitempty"`
	StudyPlanId string `bson:"studyplanid" json:"studyPlanId,omitempty"`
}

/***************************************************************/
/***************************************************************/
/* SubjectXStudyPlanResponse model for the response */
type SubjectXStudyPlanResponse struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Subject string `bson:"subject" json:"subject,omitempty"`
	StudyPlan string `bson:"studyplan" json:"studyPlan,omitempty"`
}
