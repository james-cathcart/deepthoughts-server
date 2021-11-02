package model

type Note struct {
    ID         string `json:"id"`
    Body       string `json:"body"`
    Created    int64  `json:"created"`
    Edited int64  `json:"edited"`
}
