package main

import (
	mondrv "ccovdata/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gookit/color.v1"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

func connect(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print('d')
	}

	collection = client.Database("ccov").Collection("DriverRegister")
}


func getAll() ([]*mondrv.Driver, error) {
	filter := bson.D{{}}

	return filterDriver(filter)
}

func filterDriver(filter interface{}) ([]*mondrv.Driver, error) {
	var drivers []*mondrv.Driver

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return drivers, err
	}

	for cur.Next(ctx) {
		var d mondrv.Driver
		err := cur.Decode(&d)

		if err != nil {
			fmt.Println(err)
			return drivers, err
		}

		drivers = append(drivers, &d)
	}

	if cur.Err(); err != nil {
		return drivers, err
	}

	cur.Close(ctx)

	if len(drivers) == 0 {
		return drivers, mongo.ErrNoDocuments
	}

	return drivers, nil
}

func printDrivers(drivers []*mondrv.Driver) {
	for i, v := range drivers {
		if v.BlockedDriver {
			color.Red.Printf("%d: %s\n", i+1, mondrv.BuildValidaDriver(v))
		} else {
			color.Green.Printf("%d: %s\n", i+1, mondrv.BuildValidaDriver(v))
		}
	}
}

func main()  {
	connect("mongodb://localhost:27017")

	drivers, err := getAll()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Print("Nenhum registro")
		}
	} else {
		printDrivers(drivers)
	}

	//ping(client, ctx)
}
