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
/* GetTestsDB get the all the test from db */
//func GetTestsDB() ([]*models.Test, bool) {
//	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
//	defer cancel()
//
//	db := config.MongoConnection.Database("san_martin")
//	collection := db.Collection("test")
//
//	var results []*models.Test
//
//	condition := bson.M {  }
//	optionsQuery := options.Find()
//	optionsQuery.SetSort(bson.D {{ Key: "testdate", Value: -1}})
//
//	tests, err := collection.Find(ctx, condition, optionsQuery)
//	if err != nil {
//		log.Fatal(err.Error())
//		return results, false
//	}
//
//	for tests.Next(context.TODO()) {
//		var row models.Test
//		err := tests.Decode(&row)
//		if err != nil {
//			return results, false
//		}
//		results = append(results, &row)
//	}
//
//	return results, true
//}
func GetTestsDB() ([]models.TestResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("test")

	condition := make([]bson.M, 0)

	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"subjectstudyplanid": bson.M { "$toObjectId": "$subjectstudyplanid" },
			"professorid": bson.M { "$toObjectId": "$professorid" },
			"testtypeid": bson.M { "$toObjectId": "$testtypeid" },
			"sheet": "$sheet",
			"form": "$form",
			"testdate": "$testdate",
		},
	})
	// aca conecto las dos tablas
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "subject_x_study_plan",
			"localField": "subjectstudyplanid",
			"foreignField": "_id",
			"as": "subjectxstudyplan",
	}})
	condition = append(condition, bson.M { "$unwind": "$subjectxstudyplan" })

	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "subject",
			"localField": "subjectstudyplan.subjectid",
			"foreignField": "_id",
			"as": "subject",
	}})

	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "test_type",
			"localField": "testtypeid",
			"foreignField": "_id",
			"as": "test_type",
	}})
	condition = append(condition, bson.M { "$unwind": "$test_type" })
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "professor",
			"localField": "professorid",
			"foreignField": "_id",
			"as": "professor",
	}})
	condition = append(condition, bson.M { "$unwind": "$professor" })
	// aca muestros los datos
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"test": "$test_type.type",
			"subject": "$subjectxstudyplan.subjectid",
			"professor": bson.M { "$concat": []string { "$professor.name", " ", "$professor.surname" }},
			"sheet": "$sheet",
			"form": "$form",
			"date": "$testdate",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.TestResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, 400, err
  }
	return result, 200, nil
}

/***************************************************************/
/***************************************************************/
/* InsertTestDB insert one test in db */
func InsertTestDB(t models.Test) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("test")

	row := bson.M {
		"testdate": t.TestDate,
		"sheet": t.Sheet,
		"form": t.Form,
		"subjectstudyplanid": t.SubjectStudyPlanId,
		"professorid": t.ProfessorId,
		"testtypeid": t.TestTypeId,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al insertar el examen", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* DeleteTestDB delete the test from the db */
func DeleteTestDB(IDTest string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("test")

	objID, _ := primitive.ObjectIDFromHex(IDTest)

	condition := bson.M {
		"_id": objID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}

/***************************************************************/
/***************************************************************/
/* GetTestDB get the test by id */
func GetTestDB(IDTest string) (models.Test, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("test")

	objID, _ := primitive.ObjectIDFromHex(IDTest)

	condition := bson.M {
		"_id": objID,
	}

	var test models.Test

	err := collection.FindOne(ctx, condition).Decode(&test)
	return test, err
}

