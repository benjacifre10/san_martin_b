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
/* GetStudentTestsDB get the all the test by student from db */
func GetStudentTestsDB() ([]models.StudentXTestResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student_x_test")

	condition := make([]bson.M, 0)

	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M {
			"testid": bson.M { "$toObjectId": "$testid" },
			"studentsubjectstudyplanid": bson.M { "$toObjectId": "$studentsubjectstudyplanid" },
			"test": "$test_type.type",
			"note": "$note",
//			"date": "$test.testdate",
		},
	})
	// aca conecto las dos tablas

	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "test",
			"localField": "testid",
			"foreignField": "_id",
			"as": "test",
	}})
	condition = append(condition, bson.M { "$unwind": "$test" })

	condition = append(condition, bson.M {
		"$project": bson.M {
//			"date": "$test.testdate",
			"studentsubjectstudyplanid": bson.M { "$toObjectId": "$studentsubjectstudyplanid" },
			"testid": bson.M { "$toObjectId": "$testid" },
			"testtypeid": bson.M { "$toObjectId": "$test.testtypeid" },
			"test": "$test_type.type",
			"note": "$note",
		},
	})
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "test_type",
			"localField": "testtypeid",
			"foreignField": "_id",
			"as": "test_type",
	}})
	condition = append(condition, bson.M { "$unwind": "$test_type" })

	//condition = append(condition, bson.M {
	//	"$project": bson.M {
	//		"studentsubjectstudyplanid": bson.M { "$toObjectId": "$studentsubjectstudyplanid" },
	//		"testid": bson.M { "$toObjectId": "$testid" },
	//		"testtypeid": bson.M { "$toObjectId": "$test.testtypeid" },
	//		"test": "$test_type.type",
	//		"note": "$note",
	//	},
	//})
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "student_x_subject_x_study_plan",
			"localField": "studentsubjectstudyplanid",
			"foreignField": "_id",
			"as": "studentxsubjectxstudyplan",
	}})
	condition = append(condition, bson.M { "$unwind": "$studentxsubjectxstudyplan" })

	//condition = append(condition, bson.M {
	//	"$project": bson.M {
	//		"studentsubjectstudyplanid": bson.M { "$toObjectId": "$studentsubjectstudyplanid" },
	//		"testid": bson.M { "$toObjectId": "$testid" },
	//		"testtypeid": bson.M { "$toObjectId": "$test.testtypeid" },
	//		"studentid": bson.M { "$toObjectId": "$studentxsubjectxstudyplan.studentid" },
	//		"test": "$test_type.type",
	//		"note": "$note",
	//	},
	//})
	//condition = append(condition, bson.M {
	//	"$lookup": bson.M {
	//		"from": "student",
	//		"localField": "studentxsubjectxstudyplan.studentid",
	//		"foreignField": "_id",
	//		"as": "student",
	//}})
	//condition = append(condition, bson.M { "$unwind": "$student" })

	// aca muestros los datos
	condition = append(condition, bson.M {
		"$project": bson.M {
			"date": "$testid",
			"test": "$test_type.type",
			"note": "$note",
			"student": "$studentxsubjectxstudyplan.studentid",
			"subject": "$studentxsubjectxstudyplan.subjectstudyplanid",
		},
	})

	cur, err := collection.Aggregate(ctx, condition)
	var result []models.StudentXTestResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, 400, err
  }
	return result, 200, nil
}


/***************************************************************/
/***************************************************************/
/* InsertStudentXTestDB enroll one test in db */
func InsertStudentXTestDB(s models.StudentXTest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student_x_test")

	row := bson.M {
		"testid": s.TestId,
		"note": s.Note,
		"studentsubjectstudyplanid": s.StudentSubjectStudyPlanId,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al inscribirse en el examen", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* UpdateTestNoteDB update the note from test in the db */
func UpdateTestNoteDB(s models.StudentXTest) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student_x_test")
	
	row := make(map[string]interface{})
	row["note"] = s.Note

	updateString := bson.M {
		"$set": row,
	}

	var idStudentXTest string
	idStudentXTest = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idStudentXTest)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

