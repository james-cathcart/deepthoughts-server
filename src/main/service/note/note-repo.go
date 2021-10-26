package note

import (
    "deepthoughts-server/model"
)

type NoteRepo interface {
    GetNotesForUser(userID string) []model.Note
    CreateNote(note model.Note) error
}