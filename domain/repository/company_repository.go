package repository

import (
	"ccovdata/domain/entity/ccov"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	db  *mongo.Client
	ctx context.Context
}

func NewCompanyRepository(db *mongo.Client, ctx context.Context) *CompanyRepository {
	return &CompanyRepository{
		db:  db,
		ctx: ctx,
	}
}

func (com *CompanyRepository) GetCompanyExternal(companyName string, ctx context.Context) (*ccov.Company, error) {
	var company *ccov.Company

	collection := com.db.Database("ccov").Collection("Company")

	err := collection.FindOne(ctx, bson.M{"Name": companyName}).Decode(&company)
	if err != nil {
		return company, err
	}

	return company, nil
}
