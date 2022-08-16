// Package models provides ...
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* Professor model for the mongo DB */
type Professor struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:""`
	Name string `bson:"name" json:"name,omitempty"`
	Surname string `bson:"surname" json:"surname,omitempty"`
	IdentityNumber string `bson:"identitynumber" json:"identityNumber,omitempty"`
}
