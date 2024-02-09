package app

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
    DB *mongo.Client
	HTTPServer *http.Server
}

var uri = "mongodb://root:toor@localhost:27018/admin"

func NewApp() *App {
    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatalf("Error pinging MongoDB: %v", err)
    }

    return &App{
        DB: client,
		HTTPServer: &http.Server{
			Addr: ":8080",
		},
    }
}


func (a* App) Start() {
	log.Println("Starting the server on 8080")
	
	err := a.HTTPServer.ListenAndServe()

	if err != nil {
		log.Fatalf("Error starting the server : %v", err)
	}

}