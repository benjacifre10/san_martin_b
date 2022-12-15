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
/* InsertStudentXSubjectXStudyPlanDB insert one final note of subject x study plan in db */
func InsertStudentXSubjectXStudyPlanDB(s models.StudentXSubjectXStudyPlan) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student_x_subject_x_study_plan")

	row := bson.M {
		"finalnote": s.FinalNote,
		"approved": s.Approved,
		"subjectstudyplanid": s.SubjectStudyPlanId,
		"studentid": s.StudentId,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al asociar la nota final de la materia con el plan de estudio", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* UpdateStudentXSubjectXStudyPlanDB update the final note int the subject in the db */
func UpdateStudentXSubjectXStudyPlanDB(s models.StudentXSubjectXStudyPlan) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student_x_subject_x_study_plan")

	row := make(map[string]interface{})
	row["finalnote"] = s.FinalNote
	row["approved"] = s.Approved

	updateString := bson.M {
		"$set": row,
	}

	var idStudentXSubjectXStudyPlan string
	idStudentXSubjectXStudyPlan = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idStudentXSubjectXStudyPlan)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* GetStudentXSubjectsXStudyPlanDB get the subjects of an student from db */
func GetStudentXSubjectsXStudyPlanDB(ID string) ([]models.StudentXSubjectXStudyPlanResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student_x_subject_x_study_plan")

	condition := make([]bson.M, 0)
	// aca matcheo con el id de student
	condition = append(condition, bson.M {
		"$match": bson.M { "studentid": ID },
	})
	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"subjectstudyplanid": bson.M { "$toObjectId": "$subjectstudyplanid" },
			"studentid": bson.M { "$toObjectId": "$studentid" },
			"finalnote": "$finalnote",
			"approved": "$approved",
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
			"from": "student",
			"localField": "studentid",
			"foreignField": "_id",
			"as": "student",
	}})
	condition = append(condition, bson.M { "$unwind": "$student" })
	// aca muestros los datos
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"subject": "$subjectxstudyplan.subjectid",
"student": bson.M { "$concat": []string { "$student.name", " ", "$student.surname" }},
			"finalnote": "$finalnote",
			"approved": "$approved",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.StudentXSubjectXStudyPlanResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, 400, err
  }
	return result, 200, nil
}
