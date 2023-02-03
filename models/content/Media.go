package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mediaCollection = "content_media"

type Media struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description,omitempty"`
	Keywords     []string           `bson:"keywords"`
	Size         int64              `bson:"size"`
	Extension    string             `bson:"extension"`
	MimeType     string             `bson:"mime_type"`
	OriginalName string             `bson:"original_name"`
	Path         string             `bson:"path"`
	Uploaded     primitive.M        `bson:"uploaded"`
	Updated      primitive.M        `bson:"updated,omitempty"`
}

func CreateMedia(doc *Media) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(mediaCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreateMedia: ", err)
		return bson.M{"Loc": "CreateMedia"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdateMedia(id primitive.ObjectID, doc *Media) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(mediaCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdateMedia: ", err)
		return bson.M{"Loc": "UpdateMedia"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}

func FindMedia(filter bson.M, limit ...int) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if csr, err := db.Collection(mediaCollection).Find(ctx, filter); err != nil {
		log.Fatal("FindMedia: ", err)
		return bson.M{"Loc": "FindMedia"}, err
	} else {
		var result []bson.M
		if err = csr.All(ctx, &result); err != nil {
			log.Fatal("FindMediaCursor: ", err)
			return bson.M{"Loc": "FindMediaCursor"}, err
		}
		count, err := db.Collection(mediaCollection).CountDocuments(ctx, filter)
		if err != nil {
			log.Fatal("FindMediaCount: ", err)
			return bson.M{"Loc": "FindMediaCount"}, err
		}
		return bson.M{"result": result, "count": count}, nil
	}
}
