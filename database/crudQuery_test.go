package database

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestGetCollection(t *testing.T) {
	StartConnection()

	collection := getCollection("contacts")
	assert.NotNil(t, collection)

	StopConnection()
}

func TestInsert(t *testing.T) {
	StartConnection()
	coll := getCollection("contacts")

	bsonDoc := bson.D{
		{Key: "name", Value: "contactName #1"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #1"},
		{Key: "createdAt", Value: 123455677},
	}

	result := insert(coll, bsonDoc)
	assert.NotNil(t, result.InsertedID)
	deleteMany(coll, bsonDoc)

	StopConnection()
}

func TestInsertMany(t *testing.T) {
	StartConnection()
	coll := getCollection("contacts")

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

	result := insertMany(coll, docs)
	assert.NotNil(t, result.InsertedIDs)
	assert.Equal(t, 2, len(result.InsertedIDs))
	coll.Drop(nil)

	StopConnection()
}

func TestFindAll(t *testing.T) {
	StartConnection()

	coll := getCollection("test_contacts")
	coll.Drop(nil)

	bsonDoc1 := bson.D{
		{Key: "name", Value: "contactName #1"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #1"},
		{Key: "createdAt", Value: 123455677},
	}
	insert(coll, bsonDoc1)
	bsonDoc2 := bson.D{
		{Key: "name", Value: "contactName #2"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #2"},
		{Key: "createdAt", Value: 123455677},
	}
	insert(coll, bsonDoc2)
	result := findAll(coll)
	assert.Equal(t, 2, len(result))

	coll.Drop(nil)
	result = findAll(coll)
	assert.Equal(t, 0, len(result))

	StopConnection()
}

func TestDeleteOne(t *testing.T) {
	StartConnection()

	coll := getCollection("test_contacts")
	coll.Drop(nil)

	bsonDoc1 := bson.D{
		{Key: "name", Value: "contactName #1"},
		{Key: "email", Value: "askdjasdl@mail.com"},
		{Key: "message", Value: "its a message #1"},
		{Key: "createdAt", Value: 123455677},
	}
	insert(coll, bsonDoc1)
	resultDelete := delete(coll, bson.D{
		{Key: "name", Value: "contactName #1"},
	})
	assert.Equal(t, int64(1), resultDelete.DeletedCount)
	coll.Drop(nil)

	StopConnection()
}
