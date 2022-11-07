// Package models provides ...
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* Student model for the mongo DB */
type Student struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Name string `bson:"name" json:"name,omitempty"`
	Surname string `bson:"surname" json:"surname,omitempty"`
	IdentityNumber string `bson:"identitynumber" json:"identityNumber,omitempty"`
	Address string `bson:"address" json:"address,omitempty"`
	Phone string `bson:"phone" json:"phone,omitempty"`
	Cuil string `bson:"cuil" json:"cuil,omitempty"`
	Arrears bool `bson:"arrears" json:"arrears"`
	State bool `bson:"state" json:"state"`
	UserId string `bson:"userid" json:"userId,omitempty"`
	DegreeId string `bson:"degreeid" json:"degreeId,omitempty"`
	CreatedAt time.Time `bson:"createdat" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedat" json:"updatedAt,omitempty"`
}

/***************************************************************/
/***************************************************************/
/* StudentResponse get the student info */
type StudentResponse struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Name string `bson:"name" json:"name,omitempty"`
	Surname string `bson:"surname" json:"surname,omitempty"`
	IdentityNumber string `bson:"identitynumber" json:"identityNumber,omitempty"`
	Address string `bson:"address" json:"address,omitempty"`
	Phone string `bson:"phone" json:"phone,omitempty"`
	Cuil string `bson:"cuil" json:"cuil,omitempty"`
	Arrears bool `bson:"arrears" json:"arrears"`
	State bool `bson:"state" json:"state"`
	Degree string `bson:"degree" json:"degree,omitempty"`
	User string `bson:"user" json:"user,omitempty"`
}
