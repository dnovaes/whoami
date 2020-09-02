package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection(collName string) *mongo.Collection {
	var client = Db.client
	return client.Database(database).Collection(collName)
}

func insert(coll *mongo.Collection, doc primitive.D) *mongo.InsertOneResult {
	result, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	return result
}

func insertMany(coll *mongo.Collection, docs []interface{}) *mongo.InsertManyResult {
	result, err := coll.InsertMany(context.Background(), docs)
	if err != nil {
		panic(err)
	}
	return result
}

func delete(coll *mongo.Collection, doc primitive.D) *mongo.DeleteResult {
	result, err := coll.DeleteOne(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	return result
}

func deleteMany(coll *mongo.Collection, doc primitive.D) *mongo.DeleteResult {
	result, err := coll.DeleteMany(context.Background(), doc)
	if err != nil {
		panic(err)
	}
	return result
}

func findAll(coll *mongo.Collection) []bson.M {
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
