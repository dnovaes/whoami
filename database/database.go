package database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"

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
var hostURI string

const database string = "dnovaes"

func init() {
	loadCredentials()
	hostURI = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", credentials.user, credentials.pass, credentials.host, credentials.dbName)
}

func loadCredentials() {
	viper.AddConfigPath("$HOME/Config/")
	viper.SetConfigName("portfolio")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			fmt.Println(err)
		} else {
			fmt.Println("Unable to locate Config file.", err)
			return
		}
	}

	user := viper.Get("MONGO_USER").(string)
	pass := viper.Get("MONGO_PASS").(string)
	cluster := viper.Get("MONGO_HOST").(string)
	dbName := viper.Get("MONGO_DBNAME").(string)
	credentials = DbCredentials{user, pass, cluster, dbName}
}

func StartConnection() {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(hostURI))
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
