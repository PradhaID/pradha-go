package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var AccountingCoaCollection = "accounting_coa"

type Coa struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Parent      primitive.ObjectID `bson:"parent,omitempty"`
	Code        int32              `bson:"code"`
	Name        string             `bson:"name"`
	Description string             `bson:"description,omitempty"`
	Position    string             `bson:"position"`
	Created     primitive.M        `bson:"created"`
	Updated     primitive.M        `bson:"updated,omitempty"`
}

func CreateCoa(doc *Coa) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(AccountingCoaCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreateCoa: ", err)
		return bson.M{"Loc": "CreateCoa"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdateCoa(id primitive.ObjectID, doc *Coa) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(AccountingCoaCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdateCoa: ", err)
		return bson.M{"Loc": "UpdateCoa"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}

func FindCoa(filter bson.M, sort bson.D, limit int64, page int64) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX

	opt := options.Find()
	opt.SetSort(sort)
	opt.SetLimit(limit)

	opt.SetSkip(limit * (page - 1))
	if csr, err := db.Collection(AccountingCoaCollection).Find(ctx, filter, opt); err != nil {
		log.Fatal("FindCoa: ", err)
		return bson.M{"Loc": "FindCoa"}, err
	} else {
		var result []bson.M
		if err = csr.All(ctx, &result); err != nil {
			log.Fatal("FindCoaCursor: ", err)
			return bson.M{"Loc": "FindCoaCursor"}, err
		}
		count, err := db.Collection(AccountingCoaCollection).CountDocuments(ctx, filter)
		if err != nil {
			log.Fatal("FindCoaCount: ", err)
			return bson.M{"Loc": "FindCoaCount"}, err
		}
		return bson.M{"result": result, "count": count}, nil
	}
}

func FindOneCoa(filter bson.M) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	var result bson.M
	if err := db.Collection(AccountingCoaCollection).FindOne(ctx, filter).Decode(&result); err != nil {
		log.Fatal("FindOneCoa: ", err)
		return bson.M{"Loc": "FindOneCoa"}, err
	} else {
		return result, nil
	}
}
