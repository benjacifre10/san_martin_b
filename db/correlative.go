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
/* GetCorrelativesByStudyPlanDB get the correlatives by study plan from db */
func GetCorrelativesByStudyPlanDB() ([]*models.Correlative, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("correlative")

	var result []*models.Correlative

	condition := bson.M {  }
	optionsQuery := options.Find()

	correlatives, err := collection.Find(ctx, condition, optionsQuery)
	if err != nil {
		log.Fatal(err.Error())
		return result, 400, err
	}

	for correlatives.Next(context.TODO()) {
		var row models.Correlative
		err := correlatives.Decode(&row)
		if err != nil {
			return result, 400, err
		}
		result = append(result, &row)
	}

	return result, 200, nil
}

/***************************************************************/
/***************************************************************/
/* InsertCorrelativeDB insert one correlative in db */
func InsertCorrelativeDB(c models.Correlative) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("correlative")

	row := bson.M {
		"year": c.Year,
		"correlative": c.Correlative,
		"subjectxstudyplanid": c.SubjectXStudyPlanId,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al insertar la correlatividad", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* CheckExistDegree2 check if degree already exists */
func CheckExistDegree2(nameDegree string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("degree")

	condition := bson.M {
		"name": nameDegree,
	}

	var result models.Degree

	err := collection.FindOne(ctx, condition).Decode(&result)
	if (result.Name != "") {
		return result.ID.Hex(), true, nil
	}

	return "", false, err
}

/***************************************************************/
/***************************************************************/
/* UpdateDegreeDB2 update the degree in the db */
func UpdateDegreeDB2(d models.Degree) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("degree")

	row := make(map[string]interface{})
	row["name"] = d.Name

	updateString := bson.M {
		"$set": row,
	}

	var idDegree string
	idDegree = d.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idDegree)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateStatusDegreeDB2 update the degree in the db */
func UpdateStatusDegreeDB2(d models.Degree) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("degree")

	row := make(map[string]interface{})
	row["active"] = d.Active

	updateString := bson.M {
		"$set": row,
	}

	var idDegree string
	idDegree = d.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idDegree)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
