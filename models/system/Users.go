package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var SystemUserCollection = "system_users"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Group      primitive.ObjectID `bson:"group"`
	Username   string             `bson:"username"`
	Email      string             `bson:"email"`
	Name       string             `bson:"name"`
	Password   string             `bson:"password,omitempty"`
	Phone      string             `bson:"phone,omitempty"`
	Activation string             `bson:"activation,omitempty"`
	Reset      string             `bson:"reset,omitempty"`
	Status     []primitive.M      `bson:"status"`
	Created    []primitive.M      `bson:"created"`
	Updated    []primitive.M      `bson:"updated,omitempty"`
}

func CreateUser(doc *User) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(SystemUserCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreateMedia: ", err)
		return bson.M{"Loc": "CreateMedia"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdateUser(id primitive.ObjectID, doc *User) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(SystemUserCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdateUser: ", err)
		return bson.M{"Loc": "UpdateUser"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}

func FindUser(filter bson.M, sort bson.D, limit int64) (bson.M, error) {
	opts := options.Find()
	opts.SetSort(sort)
	opts.SetLimit(limit)
	ctx := context.Background()
	db := initializers.CNX
	if csr, err := db.Collection(SystemUserCollection).Find(ctx, filter, opts); err != nil {
		log.Fatal("FindUser: ", err)
		return bson.M{"Loc": "FindUser"}, err
	} else {
		var result []bson.M
		if err = csr.All(ctx, &result); err != nil {
			log.Fatal("FindUserCursor: ", err)
			return bson.M{"Loc": "FindUserCursor"}, err
		}
		count, err := db.Collection(SystemUserCollection).CountDocuments(ctx, filter)
		if err != nil {
			log.Fatal("FindUserCount: ", err)
			return bson.M{"Loc": "FindUserCount"}, err
		}
		return bson.M{"result": result, "count": count}, nil
	}
}

func FindOneUser(filter bson.M, sort bson.D) (bson.M, error) {
	opts := options.FindOne()
	opts.SetSort(sort)
	ctx := context.Background()
	db := initializers.CNX
	var result bson.M
	if err := db.Collection(SystemUserCollection).FindOne(ctx, filter, opts).Decode(&result); err != nil {
		log.Fatal("FindOneUser: ", err)
		return bson.M{"Loc": "FindOneUser"}, err
	} else {
		return result, nil
	}
}
