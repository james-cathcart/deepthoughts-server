package note

import (
    "deepthoughts-server/model"
    "go.mongodb.org/mongo-driver/mongo"
)

type NoteRepo interface {
    GetNotesForUser(userID string) []*model.Note
    CreateNote(note model.Note) *mongo.InsertOneResult
}