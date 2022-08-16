package db

import (
	"context"
	"log"
	"time"

	"github.com/benjacifre10/san_martin_b/config"
	"github.com/benjacifre10/san_martin_b/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/***************************************************************/
/***************************************************************/
/* GetSubjectsDB get the subjects from db */
func GetSubjectsDB() ([]*models.Subject, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	var results []*models.Subject

	condition := bson.M {  }
	optionsQuery := options.Find()
	optionsQuery.SetSort(bson.D {{ Key: "name", Value: 1}})

	subjects, err := collection.Find(ctx, condition, optionsQuery)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for subjects.Next(context.TODO()) {
		var row models.Subject
		err := subjects.Decode(&row)
		if err != nil {
			return results, false
		}
		results = append(results, &row)
	}

	return results, true
}

/***************************************************************/
/***************************************************************/
/* InsertSubjectDB insert one subject in db */
func InsertSubjectDB(s models.Subject) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	row := bson.M {
		"name": s.Name,
		"professorid": s.ProfessorId,
		"shiftid": s.ShiftId,
		"pursuetypeid": s.PursueTypeId,
		"credithours": s.CreditHours,
		"days": s.Days,
		"from": s.From,
		"to": s.To,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al insertar la materia", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* CheckExistSubject check if subject already exists */
func CheckExistSubject(nameSubject string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	condition := bson.M {
		"name": nameSubject,
	}

	var result models.Subject

	err := collection.FindOne(ctx, condition).Decode(&result)
	if (result.Name != "") {
		return result.ID.Hex(), true, nil
	}

	return "", false, err
}

/***************************************************************/
/***************************************************************/
/* UpdateSubjectDB update the subject in the db */
func UpdateSubjectDB(s models.Subject) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	row := make(map[string]interface{})
	row["name"] = s.Name
	row["professorid"] = s.ProfessorId
	row["shiftid"] = s.ShiftId
	row["pursuetypeid"] = s.PursueTypeId
	row["credithours"] = s.CreditHours
	row["days"] = s.Days
	row["from"] = s.From
	row["to"] = s.To

	updateString := bson.M {
		"$set": row,
	}

	var idSubject string
	idSubject = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idSubject)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteSubjectDB delete the academy subject from the db */
func DeleteSubjectDB(IDSubject string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	objID, _ := primitive.ObjectIDFromHex(IDSubject)

	condition := bson.M {
		"_id": objID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}

/***************************************************************/
/***************************************************************/
/* GetSubjectDB get the academy subject by id */
func GetSubjectDB(IDSubject string) (models.Subject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	objID, _ := primitive.ObjectIDFromHex(IDSubject)

	condition := bson.M {
		"_id": objID,
	}

	var subject models.Subject

	err := collection.FindOne(ctx, condition).Decode(&subject)
	return subject, err
}
