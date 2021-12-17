package main

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var limit = 100

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

	for i, v := range ccovRegisterExternal {
		companyExtraData := getCompanyCnpjById(v.Company, db, ctx)
		registerExtraData := getCompleteRegisterById(v.Id, db, ctx)

		v.CompanyExtra = companyExtraData
		v.RegisterExtra = registerExtraData

		if i > limit {
			break
		}
	}
	//
	//usecase := parse_register.NewParseRegister(repo)
	//
	//usecase.Execute()
	printDrivers(ccovRegisterExternal)
}

func getCompanyCnpjById(companyName string, db *mongo.Client, ctx context.Context) *ccov.Company {
	defaultCompany := ccov.NewCompany()
	defaultCompany.CompanyIsVerttice()

	companyRepo := repository.NewCompanyRepository(db, ctx)
	companyData, err := companyRepo.GetCompanyExternal(companyName, ctx)
	if err != nil {
		return defaultCompany
	}

	return companyData
}

func getCompleteRegisterById(id primitive.ObjectID, db *mongo.Client, ctx context.Context) *ccov.DriverRegister {
	defaultRegister := ccov.NewDriverRegister()

	registerRepo := repository.NewDriverRegisterRepository(db, ctx)
	registerData, err := registerRepo.GetDriverRegister(id, ctx)

	if err != nil {
		return defaultRegister
	}

	return registerData
}

func printDrivers(drivers []*ccov.DriverRegisterExternal) {
	for i, v := range drivers {
		fmt.Printf("%d: %s\n", i+1, v.RegisterExtra.DriverProfile)

		if i > limit {
			break
		}
	}
}
