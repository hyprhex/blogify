package main

import (
	"log"
	"net/http"
)

func (app *application) internalErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s, path: %s, error: %s", r.Method, r.URL.Path, err)

	writeJSON(w, http.StatusInternalServerError, "the server encounterd a problem")
}

func (app *application) badRequestErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request response: %s, path: %s, error: %s", r.Method, r.URL.Path, err)

	writeJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found resource: %s, path: %s, error: %s", r.Method, r.URL.Path, err)

	writeJSON(w, http.StatusNotFound, "not found")
}
