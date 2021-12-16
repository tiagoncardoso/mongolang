package main

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func connect(dsn string) (error, *mongo.Client, context.Context) {
	var ctx context.Context
	clientOptions := options.Client().ApplyURI(dsn)
	client, _ := mongo.Connect(ctx, clientOptions)

	err := client.Ping(ctx, nil)
	if err != nil {
		return err, nil, nil
	}

	return nil, client, ctx
}

func main() {
	err, db, ctx := connect("mongodb://localhost:27017")

	if err != nil {
		fmt.Print("Não foi possível estabelecer a conexão com ccov via mongo db")
	}

	repo := repository.NewCcovRegisterRepository(db, ctx)
	ccovRegisterExternal, err := repo.GetRegisterExternal("DriverRegisterExternal", ctx)

	if err != nil {
		log.Fatal(err)
	}
	//
	//usecase := parse_register.NewParseRegister(repo)
	//
	//usecase.Execute()
	printDrivers(ccovRegisterExternal)
}

func printDrivers(drivers []*ccov.DriverRegisterExternal) {
	for i, v := range drivers {
		fmt.Printf("%d: %s\n", i+1, v)
	}
}
