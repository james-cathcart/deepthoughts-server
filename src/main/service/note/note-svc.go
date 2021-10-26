package note

import (
    "context"
    "deepthoughts-server/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

type NoteService struct {
    Client *mongo.Client
    collection *mongo.Collection
}

func (service *NoteService) GetNotesForUser(userID string) []*model.Note {

    findOptions := options.Find()
    findOptions.SetSort(bson.D{{"Created", -1}})
    ctx := context.TODO()
    noteCursor, err := service.getCollection().Find(ctx, bson.D{{}}, findOptions)
    if err != nil {
        panic(err)
    }
    defer noteCursor.Close(ctx)

    var results []*model.Note

    for noteCursor.Next(ctx) {
        var currentNote model.Note
        err := noteCursor.Decode(&currentNote)
        if err != nil {
            panic(err)
        }

        results = append(results, &currentNote)
    }

    log.Printf("NoteService -> retrieved: %d records from the database\n", len(results))

    if err := noteCursor.Err(); err != nil {
        panic(err)
    }

    return results
}

func (service *NoteService) CreateNote(note model.Note) *mongo.InsertOneResult {

    log.Printf("note svc -> creating new note")

    note.Created = time.Now().Unix()

    ctx := context.TODO()
    insertResult, err := service.getCollection().InsertOne(ctx, note)
    if err != nil {
        panic(err)
    }

    return insertResult
}

func (service *NoteService) UpdateNote() {

}

func (service *NoteService) DeleteNote() {

}

func (service *NoteService) getCollection() (*mongo.Collection) {
    if service.collection == nil {
        log.Println("initializing note collection..")
        service.collection = service.Client.Database("deepthoughts").Collection("note")
    }
    log.Println("returning note collection")
    return service.collection
}