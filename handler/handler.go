package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/danew/service/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func getTemplatesHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		templates, err := repo.getTemplates()
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, &templates)
		}
	}
}

func getTemplateHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		templateID := vars["id"]
		template, err := repo.getTemplate(templateID)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, &template)
		}
	}
}

func getCampaignsHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		items, err := repo.getCampaigns()
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, &items)
		}
	}
}

func getCampaignHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		campaignID := vars["id"]
		campaign, err := repo.getCampaign(campaignID)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, &campaign)
		}
	}
}

func getPublishedCampaignsHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		items, err := repo.getPublishedCampaigns()
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, &items)
		}
	}
}

func createCampaignHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newCampaignRequest newCampaignRequest
		err := json.Unmarshal(payload, &newCampaignRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse match request")
			return
		}
		if !newCampaignRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new match request")
			return
		}
		newCampaign, err := NewCampaign(newCampaignRequest)
		if err != nil {
			formatter.Text(w, http.StatusInternalServerError, "Unable to create a Campaign at this time")
			return
		}
		repo.addCampaign(newCampaign)
		w.Header().Add("Location", "/campaigns/"+newCampaign.ID)
		formatter.JSON(w, http.StatusCreated, &newCampaign)
	}
}

func publishCampaignHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		campaignID := vars["id"]
		campaign, err := repo.getCampaign(campaignID)
		campaign.Status = model.Implemented
		repo.updateCampaign(campaignID, campaign)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, &campaign)
		}
	}
}

func getResourcesHandler(formatter *render.Render, repo campaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		client := &http.Client{}
		resp, err := client.Get("https://shoelace-dev-test.azurewebsites.net/api/UserProducts")
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Unable to retrieve products")
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var products Products
		err = json.Unmarshal(body, &products)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse products request")
			return
		}
		formatter.JSON(w, http.StatusOK, &products)
	}
}
