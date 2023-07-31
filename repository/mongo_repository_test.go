package repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/charlygame/CatGameService/db"
	"github.com/charlygame/CatGameService/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func initTest() {
	test.Init("../test.env")
}

func TestFindShouldReturn0Items(t *testing.T) {
	initTest()
	defer test.Clear()

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D
	err := mongo_repository.List(nil, nil, 0, 0, nil, &results)

	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 0, len(results))
}

func TestFindShouldReturnFiltered(t *testing.T) {
	initTest()
	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "name", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D
	err := mongo_repository.List(bson.D{{Key: "id", Value: bson.D{{Key: "$gt", Value: 500}}}}, nil, 0, 0, nil, &results)

	if err != nil {
		log.Fatal(err)
	}
	doc := results[0]
	assert.Equal(t, 10, len(results))
	assert.Equal(t, int32(501), doc.Map()["id"])
}

func TestFindShouldReturnProjected(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D

	err := mongo_repository.List(nil, bson.D{{Key: "id", Value: 1}}, 0, 10, nil, &results)

	if err != nil {
		log.Fatal(err)
	}

	doc := results[0]
	assert.Equal(t, 10, len(results))
	assert.Equal(t, int32(0), doc.Map()["id"])
	assert.Equal(t, nil, doc.Map()["title"])
}

func TestFindShouldReturnLimited(t *testing.T) {
	initTest()
	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D
	mongo_repository.List(nil, nil, 0, 10, nil, &results)

	assert.Equal(t, 10, len(results))

}

func TestFindShouldReturnSkipped(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D

	err := mongo_repository.List(nil, nil, 10, 10, nil, &results)

	if err != nil {
		log.Fatal(err)
	}

	doc := results[0]

	assert.Equal(t, 10, len(results))
	assert.Equal(t, int32(10), doc.Map()["id"])
}

func TestFindShouldReturnSorted(t *testing.T) {
	initTest()
	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D

	err := mongo_repository.List(nil, nil, 0, 10, bson.D{{Key: "id", Value: -1}}, &results)

	if err != nil {
		log.Fatal(err)
	}

	doc := results[0]

	assert.Equal(t, 10, len(results))
	assert.Equal(t, int32(510), doc.Map()["id"])
}

func TestFindShouldReturn500ItemsWithLimitOverThan500(t *testing.T) {

	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 510; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var results []bson.D
	err := mongo_repository.List(nil, nil, 0, 1000, nil, &results)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 500, len(results))
}

func TestGetShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID()

	doc := bson.D{
		{Key: "_id", Value: _id},
		{Key: "title", Value: "Test Title"},
	}

	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var result bson.D

	err := mongo_repository.Get(_id.Hex(), &result)

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	log.Printf("Result: %+v", result)

	assert.Equal(t, "Test Title", result.Map()["title"])
}

func TestGetShouldFail(t *testing.T) {

	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID()

	doc := bson.D{
		{Key: "_id", Value: _id},
		{Key: "title", Value: "Test Title"},
	}

	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	var result bson.D
	id := primitive.NewObjectID().Hex()
	err := mongo_repository.Get(id, &result)

	assert.NotEqual(t, nil, err)
}

func TestCountShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	var docs bson.A

	for i := 0; i <= 500; i++ {
		doc := bson.D{
			{Key: "id", Value: int32(i)},
			{Key: "title", Value: fmt.Sprintf("Test Title %d", i)},
		}
		docs = append(docs, doc)
	}

	db.GetDB().Collection("test").InsertMany(context.Background(), docs)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	count, err := mongo_repository.Count(nil)

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	assert.Equal(t, int64(501), count)
}

func TestInsertShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	doc := bson.D{
		{Key: "title", Value: "Test Title"},
	}

	id, err := mongo_repository.Create(doc)

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	assert.NotEqual(t, nil, id)
}

func TestUpdateShouldSucceed(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "title", Value: "Test Title"},
		{Key: "_id", Value: _id},
	}

	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	update := bson.D{{
		Key: "title", Value: "Test Title Updated",
	}}

	assert.NotPanics(t, func() {
		mongo_repository.Update(_id, update)
	}, "Document not found")

}

func TestUpdateShouldRaiseError(t *testing.T) {
	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "title", Value: "Test Title"},
		{Key: "_id", Value: _id},
	}

	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	newDoc := bson.D{{
		Key: "title", Value: "Test Title Updated",
	}}

	err := mongo_repository.Update(primitive.NewObjectID().Hex(), newDoc)
	assert.NotEqual(t, nil, err)
}

func TestDeleteShouldSucceed(t *testing.T) {

	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "title", Value: "Test Title"},
		{Key: "_id", Value: _id},
	}
	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	assert.NotPanics(t, func() {
		mongo_repository.Delete(_id)
	},
		"Document not found")
}

func TestDeleteShouldRaiseError(t *testing.T) {

	initTest()

	defer test.Clear()

	_id := primitive.NewObjectID().Hex()

	doc := bson.D{
		{Key: "title", Value: "Test Title"},
		{Key: "_id", Value: _id},
	}
	db.GetDB().Collection("test").InsertOne(context.Background(), doc)

	mongo_repository := MongoRepository{
		Collection: "test",
	}

	err := mongo_repository.Delete(primitive.NewObjectID().Hex())
	assert.NotEqual(t, nil, err)
}
