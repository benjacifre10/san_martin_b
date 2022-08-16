package db

import (
	"context"
	"time"

	"github.com/benjacifre10/san_martin_b/config"
	"github.com/benjacifre10/san_martin_b/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/***************************************************************/
/***************************************************************/
/* GetStudyPlansDB get the study plans from db */
func GetStudyPlansDB() ([]models.StudyPlanResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("study_plan")

	condition := make([]bson.M, 0)

	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"degreeid": bson.M { "$toObjectId": "$degreeid" },
			"code": "$code",
			"state": "$state",
			"name": "$name",
		},
	})
	// aca conecto las dos tablas
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "degree",
			"localField": "degreeid",
			"foreignField": "_id",
			"as": "degree",
	}})
	condition = append(condition, bson.M { "$unwind": "$degree" })

	condition = append(condition, bson.M {
		"$project": bson.M { 
			"code": "$code",
			"state": "$state",
			"name": "$name",
			"degree": "$degree.name",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.StudyPlanResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, 400, err
  }
	return result, 200, nil
}

/***************************************************************/
/***************************************************************/
/* InsertStudyPlanDB insert one study plan in db */
func InsertStudyPlanDB(s models.StudyPlan) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("study_plan")

	row := bson.M {
		"name": s.Name,
		"code": s.Code,
		"state": s.State,
		"degreeid": s.DegreeId,
		"createdat": s.CreatedAt,
		"updatedat": s.UpdatedAt,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al insertar el plan de estudio", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* CheckExistStudyPlan check if study plan already exists */
func CheckExistStudyPlan(name string, code string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("study_plan")

	var condition = bson.M{}
	if code != "" {
		condition = bson.M {
			"code": code,
		}
	}

	if name != "" {
		condition = bson.M {
			"name": name,
		}
	}

	var result models.StudyPlan

	err := collection.FindOne(ctx, condition).Decode(&result)
	if (result.Name != "") {
		return result.ID.Hex(), true, nil
	}

	return "", false, err
}

/***************************************************************/
/***************************************************************/
/* UpdateStudyPlanDB update the study plan in the db */
func UpdateStudyPlanDB(s models.StudyPlan) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("study_plan")

	row := make(map[string]interface{})
	row["name"] = s.Name
	row["updatedat"] = time.Now()

	updateString := bson.M {
		"$set": row,
	}

	var idStudyPlan string
	idStudyPlan = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idStudyPlan)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* UpdateStateStudyPlanDB update the study plan in the db */
func UpdateStateStudyPlanDB(s models.StudyPlan) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("study_plan")

	row := make(map[string]interface{})
	row["state"] = s.State

	updateString := bson.M {
		"$set": row,
	}

	var idStudyPlan string
	idStudyPlan = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idStudyPlan)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
