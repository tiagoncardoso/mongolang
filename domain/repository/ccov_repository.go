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

	err = cur.Close(ctx)
	if len(drivers) == 0 {
		return drivers, mongo.ErrNoDocuments
	}

	return drivers, nil
}
