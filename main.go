package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wbluan/my-tasks/config"
	"github.com/wbluan/my-tasks/repository"
	"github.com/wbluan/my-tasks/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config.LoadConfig()

	clientOptions := options.Client().ApplyURI(config.GetMongoUri())
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)

	repository.InitMongoCollection(client)

	app := fiber.New()
	routes.SetupRoutes(app)
	port := config.GetPort()
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
