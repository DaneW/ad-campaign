package handler

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var webRoot string

// NewServer configures and returns a HTTP Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter().StrictSlash(true)

	repo := initRepository()
	initRoutes(mx, formatter, repo)
	n.UseHandler(mx)

	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo campaignRepository) {
	api := mx.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/templates", getTemplatesHandler(formatter, repo)).Methods("GET")
}

func initRepository() (repo campaignRepository) {
	repo = newInMemoryRepository()
	return
}
