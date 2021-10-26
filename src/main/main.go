package main

import (
    "context"
    "deepthoughts-server/handler"
    "deepthoughts-server/service/note"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
)

func main() {

    clientOptions := options.Client().ApplyURI("mongodb://mongo.apollo.dev:27017")
    mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        panic(err)
    }
    defer mongoClient.Disconnect(context.TODO())

    noteService := note.NoteService{
        Client: mongoClient,
    }

    noteHandler := handler.NoteHandler{
        NoteService: &noteService,
    }

    mux := http.NewServeMux()
    mux.Handle(handler.NoteApiEndpoint, &noteHandler)

    server := http.Server{
        Addr: ":8080",
        Handler: mux,
    }

    server.ListenAndServe()
    defer server.Close()
}