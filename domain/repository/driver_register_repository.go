package repository

import (
	"ccovdata/domain/entity/ccov"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRegisterRepository struct {
	db  *mongo.Client
	ctx context.Context
}

func NewDriverRegisterRepository(db *mongo.Client, ctx context.Context) *DriverRegisterRepository {
	return &DriverRegisterRepository{
		db:  db,
		ctx: ctx,
	}
}

func (com *DriverRegisterRepository) GetDriverRegister(id primitive.ObjectID, ctx context.Context) (*ccov.DriverRegister, error) {
	var register *ccov.DriverRegister

	collection := com.db.Database("ccov").Collection("DriverRegister")

	err := collection.FindOne(ctx, bson.M{"DriverExternalId": id}).Decode(&register)
	if err != nil {
		return register, err
	}

	return register, nil
}
