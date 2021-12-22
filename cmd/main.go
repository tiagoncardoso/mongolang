package main

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/repository"
	"ccovdata/usecase/parse_register"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gookit/color.v1"
	"log"
	"os"
)

var limit = 5000

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

	fmt.Println("Iniciando busca de cadastros Ccov.")

	repo := repository.NewCcovRegisterRepository(db, ctx)
	ccovRegisterExternal, err := repo.GetRegisterExternal("DriverRegisterExternal", ctx)

	fmt.Printf("Números de cadastros encontrados: %d\n\n", len(ccovRegisterExternal))

	if err != nil {
		log.Fatal(err)
	}

	var dbPortal *sql.DB
	dbPortal, err = myslqPortalConnection()
	if err != nil {
		log.Fatal(err)
	}

	var dbValida *sql.DB
	dbValida, err = mysqlValidaConnection()
	if err != nil {
		log.Fatal(err)
	}

	plus := 0
	for i, v := range ccovRegisterExternal {
		companyExtraData := getCompanyCnpjById(v.Company, db, ctx)
		registerExtraData := getCompleteRegisterById(v.Id, db, ctx)
		CompanyId := getCompanyId(companyExtraData.Document, dbPortal)

		v.CompanyExtra = companyExtraData
		v.RegisterExtra = registerExtraData
		v.CompanyPortalId = CompanyId

		newRegisterId, err := saveRegisterInValida(v, dbValida)
		if err != nil {
			log.Fatal("Não foi possível persistir o novo cadastro.")
		}

		if v.IsPlus() {
			color.Green.Printf("%d: \nPortal ID %d (%s) | %s - Plus: %t\n", i+1, v.CompanyPortalId, v.Company, v.Driver.Name, v.IsPlus())
			color.Yellow.Printf("%d\n\n", newRegisterId)
			plus += 1
		}
		color.White.Printf("%d: \nPortal ID %d (%s) | %s - Plus: %t\n", i+1, v.CompanyPortalId, v.Company, v.Driver.Name, v.IsPlus())
		color.Yellow.Printf("%d\n\n", newRegisterId)

		if i > limit {
			break
		}
	}

	fmt.Printf("Plus: %d\n", plus)

}

func myslqPortalConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_PORTAL_USER")+":"+os.Getenv("DB_PORTAL_PWD")+"@tcp("+os.Getenv("DB_PORTAL_HOST")+":"+os.Getenv("DB_PORTAL_PORT")+")/"+os.Getenv("DB_PORTAL_NAME"))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func mysqlValidaConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_VALIDA_USER")+":"+os.Getenv("DB_VALIDA_PWD")+"@tcp("+os.Getenv("DB_VALIDA_HOST")+":"+os.Getenv("DB_VALIDA_PORT")+")/"+os.Getenv("DB_VALIDA_NAME"))
	if err != nil {
		return nil, err
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

func saveRegisterInValida(register *ccov.DriverRegisterExternal, db *sql.DB) (int64, error) {
	validaRepo := repository.NewValidaDatabaseRepository(db)

	usecase := parse_register.NewParseRegister(validaRepo, register)

	driverRegisterId, err := usecase.SaveDriver()
	if err != nil {
		return 0, err
	}

	vehicles, _ := usecase.SaveVehicle()
	travel, _ := usecase.SaveTravel()

	var newRegister int64
	newRegister, err = usecase.SaveRegister(driverRegisterId, vehicles, travel)
	if err != nil {
		return -1, err
	}

	_, err = usecase.SaveResult(newRegister)
	if err != nil {
		color.Red.Printf("Erro ao tentar persistir a liberação do cadastro.")
	}

	return newRegister, nil
}
