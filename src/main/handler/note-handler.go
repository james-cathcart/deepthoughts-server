package handler

import (
    "deepthoughts-server/model"
    "deepthoughts-server/service/note"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

const (
    ActionType string = `action-type`
    CreateAction string = `create`
)

type NoteHandler struct {
    NoteService note.NoteRepo
}

func (handler *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    log.Printf("note handler -> '%s' endpoint called", NoteApiEndpoint)

    switch r.Method {
    case `POST`:
        handler.handlePost(w, r)
    case `GET`:
        handler.handleGet(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        _,_ = fmt.Fprintln(w, "method not allowed or not yet implemented")
    }

}

func (handler *NoteHandler) handlePost(w http.ResponseWriter, r *http.Request) {

    if r.Header.Get(ActionType) == CreateAction {
        log.Printf("note handler -> action: %s note\n", CreateAction)
        handler.newNote(w, r)
    }
}

func (handler *NoteHandler) newNote(w http.ResponseWriter, r *http.Request) {

    bodyBytes, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte("failed to read body"))
        return
    }

    var newNote model.Note
    err = json.Unmarshal(bodyBytes, &newNote)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte("failed to unmarshal json"))
        return
    }

    log.Printf("note handler -> saving new note: %s\n", newNote.ID)

    dbResult := handler.NoteService.CreateNote(newNote)
    if dbResult.InsertedID == nil {
        log.Printf("error -> could not persist new note to database\n")
        w.WriteHeader(http.StatusInternalServerError)
        _, _ = w.Write([]byte("could not save new note to database"))
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (handler *NoteHandler) handleGet(w http.ResponseWriter, r *http.Request) {

    log.Println("note handler -> handling GET request")

    params := r.URL.Query()

    if len(params) == 1 {
        if params.Get(`userid`) != "" {
            handler.getAllNotesForUserID(w, r)
        }
    } else {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("invalid or missing parameters"))
        return
    }

}

func (handler *NoteHandler) getAllNotesForUserID(w http.ResponseWriter, r *http.Request) {


    userID := r.URL.Query().Get("userid")
    if userID == "" {
        log.Println("note handler -> error: no 'userid' parameter was found")
    }
    notes := handler.NoteService.GetNotesForUser(userID)

    log.Printf("note handler -> getting all notes for userID: %s\n", userID)

    jsonBytes, err := json.Marshal(notes)
    if err != nil {
        panic(err)
    }

    w.WriteHeader(http.StatusOK)
    _, _ = w.Write(jsonBytes)

    return
}