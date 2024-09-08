package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hyprhex/blogify/internal/store"
)

type postKey string

const postCtx postKey = "post"

type createPostPayload struct {
	Title    string   `json:"title" validate:"required,max=100"`
	Content  string   `json:"content" validate:"required,max=1000"`
	Category string   `json:"Category" validate:"required,max=50"`
	Tags     []string `json:"tags" validate:"omitempty"`
}

type updatePostPayload struct {
	Title    *string   `json:"title" validate:"omitempty,max=100"`
	Content  *string   `json:"content" validate:"omitempty,max=1000"`
	Category *string   `json:"Category" validate:"omitempty,max=50"`
	Tags     *[]string `json:"tags" validate:"omitempty"`
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

func (app *application) listPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := app.store.Posts.List(ctx)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, posts); err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostContext(r)

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Posts.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostContext(r)

	var payload updatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestErrorResponse(w, r, err)

		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if payload.Category != nil {
		post.Category = *payload.Category
	}

	if payload.Tags != nil {
		post.Tags = *payload.Tags
	}

	if err := app.store.Posts.Update(r.Context(), post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostContext(r *http.Request) *store.Post {
	post, _ := r.Context().Value(postCtx).(*store.Post)
	return post
}
