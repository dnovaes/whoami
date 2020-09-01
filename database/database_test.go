package database

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestURI(t *testing.T) {
	expected := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", credentials.user, credentials.pass, credentials.host, credentials.dbName)
	assert.Equal(t, hostURI, expected)

	assert.True(t, len(credentials.user) > 0)
	assert.True(t, len(credentials.pass) > 0)
	assert.True(t, len(credentials.host) > 0)
	assert.True(t, len(credentials.dbName) > 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := mongo.Connect(ctx, options.Client().ApplyURI(hostURI))
	assert.Nil(t, err)
}

func TestConnection(t *testing.T) {
	setCleanUp(t)

	assert.Nil(t, Db.client)
	assert.True(t, StopConnection())

	StartConnection()
	assert.NotNil(t, Db.client)
	assert.True(t, StopConnection())
	//assert.PanicsWithError(t, "client is disconnected", func() { StopConnection() })
}

func TestIsDbConnected(t *testing.T) {
	assert.False(t, IsDbConnected())
	StartConnection()
	assert.True(t, IsDbConnected())
	StopConnection()
	assert.False(t, IsDbConnected())
}

func setCleanUp(t *testing.T) {
	t.Cleanup(func() {
		StopConnection()
	})
}
