package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SystemGroupCollection = "system_groups"

type Group struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description,omitempty"`
	Roles       []primitive.A      `bson:"roles"`
	Created     []primitive.M      `bson:"created"`
	Updated     []primitive.M      `bson:"updated,omitempty"`
}

func CreateGroup(doc *Group) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(SystemGroupCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreateGroup: ", err)
		return bson.M{"Loc": "CreateGroup"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdateGroup(id primitive.ObjectID, doc *Group) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(SystemGroupCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdateGroup: ", err)
		return bson.M{"Loc": "UpdateGroup"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}

func FindGroup(filter bson.M) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if csr, err := db.Collection(SystemGroupCollection).Find(ctx, filter); err != nil {
		log.Fatal("FindGroup: ", err)
		return bson.M{"Loc": "FindGroup"}, err
	} else {
		var result []bson.M
		if err = csr.All(ctx, &result); err != nil {
			log.Fatal("FindGroupCursor: ", err)
			return bson.M{"Loc": "FindGroupCursor"}, err
		}
		count, err := db.Collection(SystemGroupCollection).CountDocuments(ctx, filter)
		if err != nil {
			log.Fatal("FindGroupCount: ", err)
			return bson.M{"Loc": "FindGroupCount"}, err
		}
		return bson.M{"result": result, "count": count}, nil
	}
}

func FindOneGroup(filter bson.M) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	var result bson.M
	if err := db.Collection(SystemGroupCollection).FindOne(ctx, filter).Decode(&result); err != nil {
		log.Fatal("FindOneGroup: ", err)
		return bson.M{"Loc": "FindOneGroup"}, err
	} else {
		return result, nil
	}
}
