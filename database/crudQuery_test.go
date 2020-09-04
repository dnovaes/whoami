package database

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestGetCollection(t *testing.T) {
	StartConnection()

	collection := GetCollection("contacts")
	assert.NotNil(t, collection)

	StopConnection()
}

func TestInsertOne(t *testing.T) {
	StartConnection()
	coll := GetCollection("contacts")

	bsonDoc := bson.D{
		{Key: "name", Value: "contactName #1"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #1"},
		{Key: "createdAt", Value: 123455677},
	}

	result := Insert(coll, bsonDoc)
	assert.NotNil(t, result.InsertedID)
	DeleteMany(coll, bsonDoc)

	StopConnection()
}

func TestInsertMany(t *testing.T) {
	StartConnection()
	coll := GetCollection("contacts")

	docs := []interface{}{
		bson.D{
			{Key: "name", Value: "contactName #1"},
			{Key: "email", Value: "askdjasdl@mail.com"},
			{Key: "message", Value: "its a message #1"},
			{Key: "createdAt", Value: 123455677},
		},
		bson.D{
			{Key: "name", Value: "contactName #2"},
			{Key: "email", Value: "askdjasdl2@mail.com"},
			{Key: "message", Value: "its a message #2"},
			{Key: "createdAt", Value: 123455677},
		},
	}

	result := InsertMany(coll, docs)
	assert.NotNil(t, result.InsertedIDs)
	assert.Equal(t, 2, len(result.InsertedIDs))
	coll.Drop(nil)

	StopConnection()
}

func TestFindAll(t *testing.T) {
	StartConnection()

	coll := GetCollection("test_contacts")
	coll.Drop(nil)

	bsonDoc1 := bson.D{
		{Key: "name", Value: "contactName #1"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #1"},
		{"createdAt", time.Unix(1599163827, 0)},
	}
	Insert(coll, bsonDoc1)

	bsonDoc2 := bson.D{
		{Key: "name", Value: "contactName #2"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #2"},
		{"createdAt", time.Unix(1599163827, 0)},
	}
	Insert(coll, bsonDoc2)
	result := FindAll(coll)

	assert.Equal(t, 2, len(result))

	coll.Drop(nil)
	result = FindAll(coll)
	assert.Equal(t, 0, len(result))

	StopConnection()
}

func TestDeleteOne(t *testing.T) {
	StartConnection()

	coll := GetCollection("test_contacts")
	coll.Drop(nil)

	bsonDoc1 := bson.D{
		{Key: "name", Value: "contactName #1"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #1"},
		{Key: "createdAt", Value: 123455677},
	}
	Insert(coll, bsonDoc1)
	resultDelete := Delete(coll, bson.D{
		{Key: "name", Value: "contactName #1"},
	})
	assert.Equal(t, int64(1), resultDelete.DeletedCount)
	coll.Drop(nil)

	StopConnection()
}

func TestDeleteById(t *testing.T) {
	StartConnection()

	coll := GetCollection("contacts")
	bsonDoc1 := bson.D{
		{Key: "name", Value: "contactName #x"},
		{Key: "email", Value: "askdjasdlx@mail.com"},
		{Key: "message", Value: "its a message #x"},
		{Key: "createdAt", Value: 123455677},
	}
	insertedDoc := Insert(coll, bsonDoc1)

	deletedContact := FindAndDeleteContact(insertedDoc.InsertedID.(primitive.ObjectID))
	assert.Equal(t, deletedContact.ID, insertedDoc.InsertedID)

	StopConnection()
}
