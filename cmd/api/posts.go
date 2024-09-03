package main

import (
	"net/http"

	"github.com/hyprhex/blogify/internal/store"
)

type createPostPayload struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"Category"`
	Tags     []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
	}

	post := &store.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Tags:     payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
