package repository

import (
	"ccovdata/domain/entity/ccov"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CcovRegisterRepository struct {
	db  *mongo.Client
	ctx context.Context
}

func NewCcovRegisterRepository(db *mongo.Client, ctx context.Context) *CcovRegisterRepository {
	return &CcovRegisterRepository{
		db:  db,
		ctx: ctx,
	}
}

func (drr *CcovRegisterRepository) GetRegisterExternal(collectionName string, ctx context.Context) ([]*ccov.DriverRegisterExternal, error) {
	var drivers []*ccov.DriverRegisterExternal

	collection := drr.db.Database("ccov").Collection(collectionName)

	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return drivers, err
	}

	for cur.Next(ctx) {
		var dr ccov.DriverRegisterExternal

		err = cur.Decode(&dr)
		if err != nil {
			return drivers, err
		}

		drivers = append(drivers, &dr)
	}

	err = cur.Err()
	if err != nil {
		return drivers, err
	}

	_ = cur.Close(ctx)
	if len(drivers) == 0 {
		return drivers, mongo.ErrNoDocuments
	}

	return drivers, nil
}

func (drr *CcovRegisterRepository) GetRegister(collectionName string, ctx context.Context, companyName string) ([]*ccov.DriverRegister, error) {
	var drivers []*ccov.DriverRegister
	var cur *mongo.Cursor
	var err error

	collection := drr.db.Database("ccov").Collection(collectionName)

	if companyName != "" {
		cur, err = collection.Find(ctx, bson.M{"Company": companyName})
	} else {
		cur, err = collection.Find(ctx, bson.D{{}})
	}

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var dr ccov.DriverRegister

		err = cur.Decode(&dr)
		if err != nil {
			return drivers, err
		}

		drivers = append(drivers, &dr)
	}

	err = cur.Err()
	if err != nil {
		return drivers, err
	}

	err = cur.Close(ctx)
	if len(drivers) == 0 {
		return drivers, mongo.ErrNoDocuments
	}

	return drivers, nil
}
