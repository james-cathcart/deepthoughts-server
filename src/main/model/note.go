package model

type Note struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    Body       string `json:"body"`
    Created    int64  `json:"created"`
    LastViewed int64  `json:"lastViewed"`
}
