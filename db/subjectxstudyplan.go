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
/* GetSubjectsXStudyPlanDB get the subjectsstudy plans from db */
func GetSubjectsXStudyPlanDB(ID string) ([]models.SubjectXStudyPlanResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject_x_study_plan")

	condition := make([]bson.M, 0)
	// aca matcheo con el id de studyplan
	condition = append(condition, bson.M {
		"$match": bson.M { "studyplanid": ID },
	})
	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"subjectid": bson.M { "$toObjectId": "$subjectid" },
			"studyplanid": bson.M { "$toObjectId": "$studyplanid" },
		},
	})
	// aca conecto las dos tablas
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "subject",
			"localField": "subjectid",
			"foreignField": "_id",
			"as": "subject",
	}})
	condition = append(condition, bson.M { "$unwind": "$subject" })
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "study_plan",
			"localField": "studyplanid",
			"foreignField": "_id",
			"as": "study_plan",
	}})
	condition = append(condition, bson.M { "$unwind": "$study_plan" })
	// aca muestros los datos
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"subject": "$subject.name",
			"studyplan": "$study_plan.name",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.SubjectXStudyPlanResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, 400, err
  }
	return result, 200, nil
}
/***************************************************************/
/***************************************************************/
/* InsertSubjectXStudyPlanDB insert one subject x study plan in db */
func InsertSubjectXStudyPlanDB(s models.SubjectXStudyPlan) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject_x_study_plan")

	row := bson.M {
		"subjectid": s.SubjectId,
		"studyplanid": s.StudyPlanId,
	}

	result, err := collection.InsertOne(ctx, row)
	if err != nil {
		return "Hubo un error al asociar la carrera con el plan de estudio", err
	}
	
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), nil 
}

/***************************************************************/
/***************************************************************/
/* UpdateSubjectXStudyPlanDB update the subject in the study plan in the db */
func UpdateSubjectXStudyPlanDB(s models.SubjectXStudyPlan) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject_x_study_plan")

	row := make(map[string]interface{})
	row["subjectid"] = s.SubjectId

	updateString := bson.M {
		"$set": row,
	}

	var idSubjectXStudyPlan string
	idSubjectXStudyPlan = s.ID.Hex()

	objID, _ := primitive.ObjectIDFromHex(idSubjectXStudyPlan)

	filter := bson.M { "_id": bson.M { "$eq": objID }}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/***************************************************************/
/***************************************************************/
/* DeleteSubjectXStudyPlanDB delete the association from the db */
func DeleteSubjectXStudyPlanDB(IDSubjectXStudyPlan string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject_x_study_plan")

	objID, _ := primitive.ObjectIDFromHex(IDSubjectXStudyPlan)

	condition := bson.M {
		"_id": objID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}

/***************************************************************/
/***************************************************************/
/* GetSubjectXStudyPlanDB get the subject x study plan by id */
func GetSubjectXStudyPlanDB(IDSubjectXStudyPlan string) (models.SubjectXStudyPlan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject_x_study_plan")

	objID, _ := primitive.ObjectIDFromHex(IDSubjectXStudyPlan)

	condition := bson.M {
		"_id": objID,
	}

	var professor models.SubjectXStudyPlan

	err := collection.FindOne(ctx, condition).Decode(&professor)
	return professor, err
}

/***************************************************************/
/***************************************************************/
/* DeleteSubjectXStudyPlanByStudyPlanDB delete the association from the db */
func DeleteSubjectXStudyPlanByStudyPlanDB(IDSubjectXStudyPlan string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject_x_study_plan")

	condition := bson.M {
		"studyplanid": IDSubjectXStudyPlan,
	}

	_, err := collection.DeleteMany(ctx, condition)
	return err
}
