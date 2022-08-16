// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* Subject model for the mongo DB */
type Subject struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Name string `bson:"name" json:"name,omitempty"`
	ProfessorId string `bson:"professorid" json:"professorId,omitempty"`
	ShiftId string `bson:"shiftid" json:"shiftId,omitempty"`
	PursueTypeId string `bson:"pursuetypeid" json:"pursueTypeId,omitempty"`
	CreditHours int `bson:"credithours" json:"creditHours"`
	Days []string `bson:"days" json:"days,omitempty"`
	From string `bson:"from" json:"from,omitempty"`
	To string `bson:"to" json:"to,omitempty"`
}


