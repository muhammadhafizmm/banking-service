package model

type Comment struct {
	ID      *string `json:"id"`
	Title   *string `json:"title"`
	Message *string `json:"message"`
	Like    *int    `json:"like"`
}
