package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/charlygame/CatGameService/db"
	"github.com/charlygame/CatGameService/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 5
)

type MongoRepository struct {
	Collection string
}

func (r *MongoRepository) List(query interface{}, projection interface{}, skip int64, limit int64, sort interface{}, results interface{}) *utils.GameError {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	opts := options.Find()
	if sort != nil {
		opts.SetSort(sort)
	}

	if projection != nil {
		opts.SetProjection(projection)
	}

	if skip > 0 {
		opts.SetSkip(skip)
	}

	if limit > 0 {
		if limit <= 500 {
			opts.SetLimit(limit)
		} else {
			opts.SetLimit(500)
		}
	} else {
		opts.SetLimit(500)
	}

	if query == nil {
		query = bson.M{}
	}

	cursor, err := c.Find(ctx, query, opts)

	if err != nil {
		return &utils.GameError{StatusCode: 500, Err: err}
	}

	if err = cursor.All(ctx, results); err != nil {
		return &utils.GameError{StatusCode: 500, Err: err}
	}
	return nil
}

func (r *MongoRepository) FindOne(query interface{}, data interface{}) *utils.GameError {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	if query == nil {
		query = bson.M{}
	}

	if err := c.FindOne(ctx, query).Decode(data); err != nil {
		return &utils.GameError{StatusCode: 404, Err: err}
	}
	return nil
}

func (r *MongoRepository) Get(id string, result interface{}) *utils.GameError {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	objectId, objectIdParseError := primitive.ObjectIDFromHex(id)
	fmt.Printf("objectId: %v\n", objectId)
	if objectIdParseError == nil {
		if err := c.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(result); err != nil {
			fmt.Printf("err: %v\n", err)
			return &utils.GameError{StatusCode: 404, Err: err}
		}
	} else {
		if err := c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result); err != nil {
			return &utils.GameError{StatusCode: 404, Err: err}
		}
	}
	return nil
}

func (r *MongoRepository) Count(query interface{}) (int64, *utils.GameError) {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	if query == nil {
		query = bson.M{}
	}

	count, err := c.CountDocuments(ctx, query)
	if err != nil {
		return 0, &utils.GameError{StatusCode: 500, Err: err}
	}

	return count, nil

}

func (r *MongoRepository) Create(document interface{}) (interface{}, *utils.GameError) {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	insertResult, err := c.InsertOne(ctx, document)
	if err != nil {
		return nil, &utils.GameError{StatusCode: 500, Err: err}
	}

	return insertResult.InsertedID, nil
}

func (r *MongoRepository) Update(id string, document interface{}) *utils.GameError {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	objectId, objectIdParseError := primitive.ObjectIDFromHex(id)

	var (
		updateResult *mongo.UpdateResult
		err          error
	)

	if objectIdParseError == nil {
		updateResult, err = c.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: document}})
	} else {
		updateResult, err = c.UpdateOne(ctx, bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: document}})
	}

	if err != nil {
		return &utils.GameError{StatusCode: 500, Err: err}
	}

	if updateResult.MatchedCount == 0 {
		return &utils.GameError{StatusCode: 404, Err: errors.New("document not found")}
	}

	return nil
}

func (r *MongoRepository) Delete(id string) *utils.GameError {
	c := db.GetDB().Collection(r.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	var (
		deleteResult *mongo.DeleteResult
		err          error
	)

	objectId, objectIdParseError := primitive.ObjectIDFromHex(id)

	if objectIdParseError == nil {
		deleteResult, err = c.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectId}})
	} else {
		deleteResult, err = c.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	}

	if err != nil {
		return &utils.GameError{StatusCode: 500, Err: err}
	}

	if deleteResult.DeletedCount == 0 {
		return &utils.GameError{StatusCode: 404, Err: errors.New("document not found")}
	}
	return nil
}
