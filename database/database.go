package database

import (
	"context"
	"fmt"
  "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbConnection struct {
	ctx    context.Context
	Cancel context.CancelFunc
	client *mongo.Client
}

type DbCredentials struct {
	user   string
	pass   string
	host   string
	dbName string
}

var Db DbConnection
var credentials DbCredentials

func init() {
	loadCredentials()
}

func loadCredentials() {
	user := os.Getenv("MONGO_USERNAME")
	pass := ""
	clusterHost := os.Getenv("MONGO_CLUSTER_HOST")
	dbName := ""
	credentials = DbCredentials{user, pass, clusterHost, dbName}
}

func StartConnection() {
	ctx, cancel := context.WithCancel(context.Background())

  hostURICertificate := fmt.Sprintf(
    "mongodb+srv://%s/?authSource=%s&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=%s",
    credentials.host,
    "%24external",
    "/Users/diego.novaes/CredentialsConfig/whoami-cert.pem",
  )
  fmt.Println(hostURICertificate)

  /*
  var hostURIPassword string = fmt.Sprintf(
    "mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
    credentials.user,
    credentials.pass,
    credentials.host,
    credentials.dbName,
  )
  */

  opts := options.Client().ApplyURI(hostURICertificate)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	Db = DbConnection{ctx, cancel, client}
}

func StopConnection() bool {
	if Db.client == nil {
		return true
	}

	err := Db.client.Disconnect(Db.ctx)
	if err != nil {
		panic(err)
		return false
	}
	Db.Cancel()
	Db.client = nil
	return true
}

func IsDbConnected() bool {
	if Db.client == nil {
		return false
	}
	err := Db.client.Ping(Db.ctx, readpref.Primary())
	if err != nil {
		return false
	}
	return true
}
