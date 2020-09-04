package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"github.com/dnovaes/portfolio/gqlgen/graph/model"
)

func GetCollection(collName string) *mongo.Collection {
	var client = Db.client
	return client.Database(database).Collection(collName)
}

func Insert(coll *mongo.Collection, doc primitive.D) *mongo.InsertOneResult {
	result, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	return result
}

func InsertMany(coll *mongo.Collection, docs []interface{}) *mongo.InsertManyResult {
	result, err := coll.InsertMany(context.Background(), docs)
	if err != nil {
		panic(err)
	}
	return result
}

func Delete(coll *mongo.Collection, doc primitive.D) *mongo.DeleteResult {
	result, err := coll.DeleteOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	return result
}

func DeleteMany(coll *mongo.Collection, doc primitive.D) *mongo.DeleteResult {
	result, err := coll.DeleteMany(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	return result
}

func FindAll(coll *mongo.Collection) []bson.M {
	var results []bson.M
	cursorResult, err := coll.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}
	err = cursorResult.All(context.TODO(), &results)
	if err != nil {
		panic(err)
	}

	return results
}

func FindAllContacts() []*model.Contact {
	coll := GetCollection("contacts")
	var results []*model.Contact
	cursorResult, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	err = cursorResult.All(context.TODO(), &results)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return results
}

func FindAndDeleteContact(id primitive.ObjectID) *model.Contact {
	coll := GetCollection("contacts")

	var deletedContact *model.Contact
	err := coll.FindOneAndDelete(context.Background(), bson.D{{"_id", id}}).Decode(&deletedContact)
	if err == mongo.ErrNoDocuments {
		log.Println("FindAndDeleteContact: ", err)
		return nil
	}
	return deletedContact
}
