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
/* GetProfessorsDB get the professors from db */
func GetProfessorsDB() ([]*models.Professor, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("professor")

	var results []*models.Professor

	condition := bson.M {  }
	optionsQuery := options.Find()
	optionsQuery.SetSort(bson.D {{ Key: "identitynumber", Value: 1}})

	professors, err := collection.Find(ctx, condition, optionsQuery)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for professors.Next(context.TODO()) {
		var row models.Professor
		err := professors.Decode(&row)
		if err != nil {
			return results, false
		}
		results = append(results, &row)
	}

	return results, true
}

/***************************************************************/
/***************************************************************/
/* InsertProfessorDB insert one professor in db */
func InsertProfessorDB(p models.Professor) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("professor")

	row := bson.M {
		"name": p.Name,
		"surname": p.Surname,
		"identitynumber": p.IdentityNumber,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al insertar el profesor", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* CheckExistProfessor check if professor already exists */
func CheckExistProfessor(identityNumber string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("professor")

	condition := bson.M {
		"identitynumber": identityNumber,
	}

	var result models.Professor

	err := collection.FindOne(ctx, condition).Decode(&result)
	if (result.IdentityNumber != "") {
		return result.ID.Hex(), true, nil
	}

	return "", false, err
}

/***************************************************************/
/***************************************************************/
/* UpdateProfessorDB update the professor in the db */
func UpdateProfessorDB(p models.Professor) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("professor")

	row := make(map[string]interface{})
	row["name"] = p.Name
	row["surname"] = p.Surname

	updateString := bson.M {
		"$set": row,
	}

	var idProfessor string
	idProfessor = p.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idProfessor)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteProfessorDB delete the professor from the db */
func DeleteProfessorDB(IDProfessor string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("professor")

	objID, _ := primitive.ObjectIDFromHex(IDProfessor)

	condition := bson.M {
		"_id": objID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}

/***************************************************************/
/***************************************************************/
/* GetProfessorDB get the professor by id */
func GetProfessorDB(IDProfessor string) (models.Professor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("professor")

	objID, _ := primitive.ObjectIDFromHex(IDProfessor)

	condition := bson.M {
		"_id": objID,
	}

	var professor models.Professor

	err := collection.FindOne(ctx, condition).Decode(&professor)
	return professor, err
}

