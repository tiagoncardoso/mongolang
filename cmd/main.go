package main

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var limit = 1000

func connect(dsn string) (error, *mongo.Client, context.Context) {
	var ctx context.Context

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI(dsn)
	client, _ := mongo.Connect(ctx, clientOptions)

	err = client.Ping(ctx, nil)
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

	var dbPortal *sql.DB
	dbPortal, err = myslqPortalConnection()
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range ccovRegisterExternal {
		companyExtraData := getCompanyCnpjById(v.Company, db, ctx)
		registerExtraData := getCompleteRegisterById(v.Id, db, ctx)
		CompanyId := getCompanyId(companyExtraData.Document, dbPortal)

		v.CompanyExtra = companyExtraData
		v.RegisterExtra = registerExtraData
		v.CompanyPortalId = CompanyId

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

func myslqPortalConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_PORTAL_USER")+":"+os.Getenv("DB_PORTAL_PWD")+"@tcp("+os.Getenv("DB_PORTAL_HOST")+":"+os.Getenv("DB_PORTAL_PORT")+")/"+os.Getenv("DB_PORTAL_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
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

func getCompanyId(cnpj string, db *sql.DB) int {
	portalRepo := repository.NewPortalDatabaseRepository(db)

	companyId, err := portalRepo.FindCompanyId(cnpj)
	if err != nil {
		companyId = 38
	}

	return companyId
}

func printDrivers(drivers []*ccov.DriverRegisterExternal) {
	for i, v := range drivers {
		fmt.Printf("%d: %d\n", i+1, v.CompanyPortalId)

		if i > limit {
			break
		}
	}
}
