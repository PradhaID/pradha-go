package models

import (
	"context"
	"log"
	"pradha-go/initializers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AccountingAccountCollection = "accounting_accounts"

type Account struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Coa           primitive.ObjectID   `bson:"coa"`
	Number        int64                `bson:"number"`
	Name          string               `bson:"name"`
	Description   string               `bson:"description,omitempty"`
	Balance       primitive.Decimal128 `bson:"balance"`
	LockedBalance primitive.Decimal128 `bson:"locked_balance,omitempty"`
	Status        string               `bson:"status"`
	Created       primitive.M          `bson:"created"`
	Updated       primitive.M          `bson:"updated,omitempty"`
	Confirmed     primitive.M          `bson:"confirmed,omitempty"`
}

func CreateAccount(doc *Account) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if insert, err := db.Collection(AccountingAccountCollection).InsertOne(ctx, &doc); err != nil {
		log.Fatal("CreateAccount: ", err)
		return bson.M{"Loc": "CreateAccount"}, err
	} else {
		return bson.M{"id": insert.InsertedID}, nil
	}
}

func UpdateAccount(id primitive.ObjectID, doc *Account) (bson.M, error) {
	ctx := context.Background()
	db := initializers.CNX
	if update, err := db.Collection(AccountingAccountCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": &doc}); err != nil {
		log.Fatal("UpdateAccount: ", err)
		return bson.M{"Loc": "UpdateAccount"}, err
	} else {
		return bson.M{"MatchedCount": update.MatchedCount, "ModifiedCount": update.ModifiedCount, "UpsertedID": update.UpsertedID, "UpsertedCount": update.UpsertedCount}, nil
	}
}
