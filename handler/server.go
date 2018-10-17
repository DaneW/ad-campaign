package handler

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

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
	api.HandleFunc("/templates/{id}", getTemplateHandler(formatter, repo)).Methods("GET")

	api.HandleFunc("/campaigns", getCampaignsHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/campaigns/create", createCampaignHandler(formatter, repo)).Methods("POST")
	api.HandleFunc("/campaigns/published", getPublishedCampaignsHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/campaigns/{id}", getCampaignHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/campaigns/{id}/publish", publishCampaignHandler(formatter, repo)).Methods("PUT")
}

func initRepository() (repo campaignRepository) {
	repo = newInMemoryRepository()
	return
}
