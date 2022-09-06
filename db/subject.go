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
/* GetSubjectsDB get the subjects from db */
func GetSubjectsDB() ([]models.SubjectResponse, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	condition := make([]bson.M, 0)

	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"professorid": bson.M { "$toObjectId": "$professorid" },
			"shiftid": bson.M { "$toObjectId": "$shiftid" },
			"pursuetypeid": bson.M { "$toObjectId": "$pursuetypeid" },
			"name": "$name",
			"credithours": "$credithours",
			"days": "$days",
			"from": "$from",
			"to": "$to",
		},
	})
	// aca conecto las dos tablas
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "professor",
			"localField": "professorid",
			"foreignField": "_id",
			"as": "professor",
	}})
	condition = append(condition, bson.M { "$unwind": "$professor" })
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "shift",
			"localField": "shiftid",
			"foreignField": "_id",
			"as": "shift",
	}})
	condition = append(condition, bson.M { "$unwind": "$shift" })
	condition = append(condition, bson.M {
		"$lookup": bson.M {
			"from": "pursue_type",
			"localField": "pursuetypeid",
			"foreignField": "_id",
			"as": "pursuetype",
	}})
	condition = append(condition, bson.M { "$unwind": "$pursuetype" })

	condition = append(condition, bson.M {
		"$project": bson.M { 
			"name": "$name",
			"professor": bson.M { "$concat": []string {"$professor.name", " ", "$professor.surname"} },
			"shift": "$shift.type",
			"pursuetype": "$pursuetype.type",
			"credithours": "$credithours",
			"days": "$days",
			"from": "$from",
			"to": "$to",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.SubjectResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, 400, err
  }
	return result, 200, nil
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
func CheckExistSubject(idSubject string, nameSubject string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("subject")

	condition := bson.M {
		"name": nameSubject,
	}

	var result models.Subject

	err := collection.FindOne(ctx, condition).Decode(&result)
	if (idSubject == "") {
		if (result.Name != "") {
			return result.ID.Hex(), true, nil
		}
	} else {
		if (result.Name != "" && idSubject != result.ID.Hex()) {
			return result.ID.Hex(), true, nil
		}
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
