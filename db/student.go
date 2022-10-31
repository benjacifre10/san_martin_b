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
/* InsertStudentDB insert one student in db */
func InsertStudentDB(s models.Student) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	row := bson.M {
		"name": s.Name,
		"surname": s.Surname,
		"identitynumber": s.IdentityNumber,
		"address": s.Address,
		"phone": s.Phone,
		"cuil": s.Cuil,
		"arrears": s.Arrears,
		"state": s.State,
		"userid": s.UserId,
		"degreeid": s.DegreeId,
		"createdat": s.CreatedAt,
		"updatedat": s.UpdatedAt,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al insertar el alumno", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* CheckExistStudent check if student already exists */
func CheckExistStudent(identityNumber string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	condition := bson.M {
		"identitynumber": identityNumber,
	}

	var result models.Student

	err := collection.FindOne(ctx, condition).Decode(&result)
	if (result.IdentityNumber != "") {
		return result.ID.Hex(), true, nil
	}

	return "", false, err
}

/***************************************************************/
/***************************************************************/
/* GetStudentsDB get the students from db */
func GetStudentsDB() ([]*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	var results []*models.Student

	condition := bson.M {  }
	optionsQuery := options.Find()
	optionsQuery.SetSort(bson.D {{ Key: "identitynumber", Value: 1}})

	students, err := collection.Find(ctx, condition, optionsQuery)
	if err != nil {
		log.Fatal(err.Error())
		return results, err
	}

	for students.Next(context.TODO()) {
		var row models.Student
		err := students.Decode(&row)
		if err != nil {
			return results, err
		}
		results = append(results, &row)
	}

	return results, nil
}

/***************************************************************/
/***************************************************************/
/* GetStudentByIdDB get the student by id */
func GetStudentByIdDB(IDStudent string) (models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	objID, _ := primitive.ObjectIDFromHex(IDStudent)

	condition := bson.M {
		"_id": objID,
	}

	var student models.Student

	err := collection.FindOne(ctx, condition).Decode(&student)
	return student, err
}

/***************************************************************/
/***************************************************************/
/* UpdateStudentDB update the student in the db */
func UpdateStudentDB(s models.Student) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	row := make(map[string]interface{})
	row["name"] = s.Name
	row["surname"] = s.Surname
	row["address"] = s.Address
	row["phone"] = s.Phone
	row["cuil"] = s.Cuil

	updateString := bson.M {
		"$set": row,
	}

	var idStudent string
	idStudent = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idStudent)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateStatusStudentDB update the student in the db */
func UpdateStatusStudentDB(s models.Student) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	row := make(map[string]interface{})
	row["state"] = s.State

	updateString := bson.M {
		"$set": row,
	}

	var idStudent string
	idStudent = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idStudent)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
