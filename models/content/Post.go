package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var postCollection = "content_posts"

type Post struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty"`
	Title            string               `bson:"title"`
	Content          string               `bson:"content,omitempty"`
	Description      string               `bson:"description,omitempty"`
	Tags             []string             `bson:"tags"`
	Categories       []primitive.ObjectID `bson:"categories"`
	Language         string               `bson:"language"`
	Thumbnail        string               `bson:"thumbnail"`
	ThumbnailCaption string               `bson:"thumbnail_caption"`
	Slug             string               `bson:"slug"`
	Status           string               `bson:"status"`
	Comment          bool                 `bson:"comment,omitempty"`
	Created          primitive.M          `bson:"created"`
	Updated          primitive.M          `bson:"updated,omitempty"`
	Published        primitive.M          `bson:"published,omitempty"`
}

func CreatePost(doc *Post) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(postCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreatePost: ", err)
		return bson.M{"Loc": "InsertPost"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdatePost(id primitive.ObjectID, doc *Post) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(postCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdatePost: ", err)
		return bson.M{"Loc": "UpdatePost"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}

func FindPost(filter bson.M, limit ...int) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if csr, err := db.Collection(postCollection).Find(ctx, filter); err != nil {
		log.Fatal("FindPost: ", err)
		return bson.M{"Loc": "FindPost"}, err
	} else {
		var result []bson.M
		if err = csr.All(ctx, &result); err != nil {
			log.Fatal("FindPostCursor: ", err)
			return bson.M{"Loc": "FindPostCursor"}, err
		}
		count, err := db.Collection(postCollection).CountDocuments(ctx, filter)
		if err != nil {
			log.Fatal("FindPostCount: ", err)
			return bson.M{"Loc": "FindPostCount"}, err
		}
		return bson.M{"result": result, "count": count}, nil
	}
}
