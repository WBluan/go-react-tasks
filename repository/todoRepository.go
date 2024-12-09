package repository

import (
	"context"

	"github.com/wbluan/my-tasks/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitMongoCollection(client *mongo.Client) {
	collection = client.Database("golang_db").Collection("todos")
}

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func CreateTodos(todo *models.Todo) (*models.Todo, error) {
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func UpdateTodoByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func DeleteTodoByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}
