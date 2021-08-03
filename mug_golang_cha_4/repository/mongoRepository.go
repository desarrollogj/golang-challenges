package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/desarrollogj/mug_golang_cha_4/model"
)

const DATABASE_URI string = "mongodb://localhost:27017"
const DATABASE_NAME string = "example"
const COLLECTION_NAME string = "Appointments"

type MongoRepository struct {
}

func (m MongoRepository) FindAll() []model.Appointment {
	ctx, client := connect()

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	cursor, _ := collection.Find(ctx, bson.M{})
	items := []model.Appointment{}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result model.Appointment
		cursor.Decode(&result)
		items = append(items, result)
	}

	disconnect(ctx, client)

	return items
}

func (m MongoRepository) Find(id int) model.Appointment {
	ctx, client := connect()

	result := model.Appointment{}
	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	collection.FindOne(ctx, model.Appointment{Id: id}).Decode(&result)

	disconnect(ctx, client)

	return result
}

func (m MongoRepository) Create(item model.Appointment) model.Appointment {
	ctx, client := connect()

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	collection.InsertOne(ctx, item)

	disconnect(ctx, client)

	return item
}

func (m MongoRepository) Update(item model.Appointment) model.Appointment {
	update := bson.M{"$set": bson.M{"title": item.Title, "isDone": item.IsDone}}

	ctx, client := connect()

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	collection.UpdateOne(ctx, model.Appointment{Id: item.Id}, update)

	disconnect(ctx, client)

	return item
}

func (m MongoRepository) Delete(id int) {
	ctx, client := connect()

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	collection.DeleteMany(ctx, model.Appointment{Id: id})

	disconnect(ctx, client)
}

func connect() (context.Context, *mongo.Client) {
	ctx := context.Background()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URI))

	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
