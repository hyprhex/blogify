package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hyprhex/blogify/internal/store"
)

type createPostPayload struct {
	Title    string   `json:"title" validate:"required,max=100"`
	Content  string   `json:"content" validate:"required,max=1000"`
	Category string   `json:"Category" validate:"required,max=50"`
	Tags     []string `json:"tags" validate:"omitempty"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}

	post := &store.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Tags:     payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}

		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}
