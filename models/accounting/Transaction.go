package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AccountingTransactionCollection = "accounting_transactions"

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Type        string             `bson:"type,omitempty"`
	Code        int32              `bson:"code"`
	Reference   string             `bson:"reference,omitempty"`
	Information string             `bson:"information,omitempty"`
	Journals    primitive.A        `bson:"journals"`
	Status      string             `bson:"status"`
	Created     primitive.M        `bson:"created"`
	Updated     primitive.M        `bson:"updated,omitempty"`
	Confirmed   primitive.M        `bson:"confirmed,omitempty"`
}

func CreateTransaction(doc *Transaction) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(AccountingCoaCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreateTransaction: ", err)
		return bson.M{"Loc": "CreateTransaction"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdateTransaction(id primitive.ObjectID, doc *Transaction) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(AccountingCoaCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdateTransaction: ", err)
		return bson.M{"Loc": "UpdateTransaction"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}
