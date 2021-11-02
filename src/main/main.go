package main

import (
    "context"
    "deepthoughts-server/handler"
    "deepthoughts-server/service/note"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
)

func main() {

    fmt.Println("DeepThoughts Server")
    fmt.Println("loading...")

    clientOptions := options.Client().ApplyURI("mongodb://mongo.apollo.dev:27017")
    ctx := context.TODO()
    mongoClient, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        panic(err)
    }
    defer mongoClient.Disconnect(ctx)

    noteService := note.NoteService{
        Client: mongoClient,
    }

    noteHandler := handler.NoteHandler{
        NoteService: &noteService,
    }

    mux := http.NewServeMux()
    mux.Handle(handler.NoteApiEndpoint, handler.DisableCors(&noteHandler))

    server := http.Server{
        Addr: ":8080",
        Handler: mux,
    }

    fmt.Println("starting server")
    server.ListenAndServe()
}