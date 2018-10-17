package handler

import (
	"net/http"

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
