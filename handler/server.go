package handler

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/danew/service/model"
	"github.com/google/uuid"
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
	webRoot = os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		root, err := os.Getwd()
		if err != nil {
			panic("Could not retrieve working directory")
		} else {
			webRoot = root
		}
	}

	api := mx.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/products", getResourcesHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/templates", getTemplatesHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/templates/{id}", getTemplateHandler(formatter, repo)).Methods("GET")

	api.HandleFunc("/campaigns", getCampaignsHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/campaigns/create", createCampaignHandler(formatter, repo)).Methods("POST")
	api.HandleFunc("/campaigns/published", getPublishedCampaignsHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/campaigns/{id}", getCampaignHandler(formatter, repo)).Methods("GET")
	api.HandleFunc("/campaigns/{id}/publish", publishCampaignHandler(formatter, repo)).Methods("PUT")

	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/build")))
}

func initRepository() (repo campaignRepository) {
	repo = newInMemoryRepository()
	createTemplates(repo)
	return
}

func createTemplates(repo campaignRepository) {
	adID, _ := uuid.NewRandom()
	ad := model.Ad{
		ID:    adID.String(),
		Title: "Default Title",
		Copy:  "Default Text",
	}
	templateID, _ := uuid.NewRandom()
	template := model.Template{
		ID:               templateID.String(),
		Title:            "Single Image Ad",
		Ads:              []model.Ad{ad},
		CompainObjective: model.LeadGeneration,
	}
	repo.addTemplate(template)

	adID, _ = uuid.NewRandom()
	adOne := model.Ad{
		ID:    adID.String(),
		Title: "Default Title1",
		Copy:  "Default Text1",
	}
	adID, _ = uuid.NewRandom()
	adTwo := model.Ad{
		ID:    adID.String(),
		Title: "Default Title2",
		Copy:  "Default Text2",
	}
	adID, _ = uuid.NewRandom()
	adThree := model.Ad{
		ID:    adID.String(),
		Title: "Default Title3",
		Copy:  "Default Text3",
	}
	templateID, _ = uuid.NewRandom()
	template = model.Template{
		ID:               templateID.String(),
		Title:            "Multi Image Carousel Ad",
		Ads:              []model.Ad{adOne, adTwo, adThree},
		CompainObjective: model.Conversions,
	}
	repo.addTemplate(template)

	adID, _ = uuid.NewRandom()
	adOne = model.Ad{
		ID:    adID.String(),
		Title: "Default Title1",
		Copy:  "Default Text1",
	}
	adID, _ = uuid.NewRandom()
	adTwo = model.Ad{
		ID:    adID.String(),
		Title: "Default Title2",
		Copy:  "Default Text2",
	}
	adID, _ = uuid.NewRandom()
	adThree = model.Ad{
		ID:    adID.String(),
		Title: "Default Title3",
		Copy:  "Default Text3",
	}
	templateID, _ = uuid.NewRandom()
	template = model.Template{
		ID:               templateID.String(),
		Title:            "Multi Image Slider Ad",
		Ads:              []model.Ad{adOne, adTwo, adThree},
		CompainObjective: model.Impressions,
	}
	repo.addTemplate(template)
}
