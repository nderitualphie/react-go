package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TODO struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BODY      string             `json:"body"`
	COMPLETED bool               `json:"completed"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTED TO MONGODB ATLAS")
	collection = client.Database("react_go").Collection("todos")
	e := echo.New()
	e.GET("/api/todo", getTodos)
	e.POST("/api/todo", createTodos)
	e.PATCH("/api/todo/:id", updateTodo)
	e.DELETE("/api/todo/:id", deleteTodo)
	port := os.Getenv("PORT")
	log.Fatal(e.Start("0.0.0.0:" + port))
}
func getTodos(c echo.Context) error {
	var todos []TODO
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo TODO
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(200, todos)
}
func createTodos(c echo.Context) error {
	todo := new(TODO)
	if err := c.Bind(&todo); err != nil {
		return err
	}
	if todo.BODY == "" {
		return c.String(400, "body cannot be empty")
	}
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.JSON(201, todo)
}
func updateTodo(c echo.Context) error {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.String(400, "invalid  todo id")
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.String(200, "success")
}
func deleteTodo(c echo.Context) error {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.String(400, "invalid  todo id")
	}
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.String(200, "success")
}
