package handler

import (
	"github.com/danew/service/model"
)

type campaignRepository interface {
	getTemplates() (templates []model.Template, err error)
}
