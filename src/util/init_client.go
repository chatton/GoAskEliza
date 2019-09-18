package util

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func InitClient() *mongo.Client {
	connectionStr, ok := os.LookupEnv("CONNECTION_STRING")
	if !ok {
		fmt.Println("no CONNECTION_STRING provided")
		os.Exit(1)
	}

	fmt.Printf("using connection string [%s]", connectionStr)

	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))


	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}
