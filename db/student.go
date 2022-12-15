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
func GetStudentsDB() ([]models.StudentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	condition := make([]bson.M, 0)	

	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"degreeid": bson.M { "$toObjectId": "$degreeid" },
			"userid": bson.M { "$toObjectId": "$userid" },
			"name": "$name",
			"surname": "$surname",
			"identitynumber": "$identitynumber",
			"address": "$address",
			"phone": "$phone",
			"cuil": "$cuil",
			"arrears": "$arrears",
			"state": "$state",
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
		"$lookup": bson.M {
			"from": "user",
			"localField": "userid",
			"foreignField": "_id",
			"as": "user",
	}})
	condition = append(condition, bson.M { "$unwind": "$user" })

	condition = append(condition, bson.M {
		"$project": bson.M { 
			"name": "$name",
			"surname": "$surname",
			"degree": "$degree.name",
			"user": "$user.email",
			"identitynumber": "$identitynumber",
			"address": "$address",
			"phone": "$phone",
			"cuil": "$cuil",
			"arrears": "$arrears",
			"state": "$state",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.StudentResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, err
  }
	return result, nil
}

/***************************************************************/
/***************************************************************/
/* GetStudentByIdUserDB get the student by email */
func GetStudentByIdUserDB(EmailParam string) ([]models.StudentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoConnection.Database("san_martin")
	collection := db.Collection("student")

	condition := make([]bson.M, 0)	

	// project me sirve tanto para dejar afuera a algunos campos
	// como tambien para calcular y mostrar otros
	condition = append(condition, bson.M {
		"$project": bson.M { 
			"degreeid": bson.M { "$toObjectId": "$degreeid" },
			"userid": bson.M { "$toObjectId": "$userid" },
			"name": "$name",
			"surname": "$surname",
			"identitynumber": "$identitynumber",
			"address": "$address",
			"phone": "$phone",
			"cuil": "$cuil",
			"arrears": "$arrears",
			"state": "$state",
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
		"$lookup": bson.M {
			"from": "user",
			"localField": "userid",
			"foreignField": "_id",
			"as": "user",
	}})
	condition = append(condition, bson.M { "$unwind": "$user" })
	condition = append(condition, bson.M { "$match": bson.M { "user.email": EmailParam }})

	condition = append(condition, bson.M {
		"$project": bson.M { 
			"name": "$name",
			"surname": "$surname",
			"degree": "$degree.name",
			"user": "$user.email",
			"identitynumber": "$identitynumber",
			"address": "$address",
			"phone": "$phone",
			"cuil": "$cuil",
			"arrears": "$arrears",
			"state": "$state",
		},
	})
	
	cur, err := collection.Aggregate(ctx, condition)
	var result []models.StudentResponse

	err = cur.All(ctx, &result)
	if err != nil {
		return result, err
  }
	return result, nil

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
	row["updatedat"] = s.UpdatedAt

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
