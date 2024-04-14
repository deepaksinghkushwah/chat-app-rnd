package main

import (
	"context"
    "time"
	"fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Example app to read data from kafka and store in monogodb")

	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	defer cancel();

	client, err := mongo.connect(ctx, options.Client(), ApplyURI("mongodb://0.0.0.0:27017"))

	defer func(){
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
}
