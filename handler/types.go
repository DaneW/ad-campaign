package handler

import (
	"github.com/danew/ad-campaign/model"
	"github.com/google/uuid"
)

type campaignRepository interface {
	addCampaign(campaign model.Campaign) (err error)
	getCampaign(id string) (campaign model.Campaign, err error)
	getCampaigns() (campaigns []model.Campaign, err error)
	getPublishedCampaigns() (campaigns []model.Campaign, err error)
	updateCampaign(id string, campaign model.Campaign) (err error)
	addTemplate(template model.Template) (err error)
	getTemplate(id string) (template model.Template, err error)
	getTemplates() (templates []model.Template, err error)
}

// NewCampaignRequest represents the request body to create a Campaign
type newCampaignRequest struct {
	Template  string   `json:"template"`
	AdTitle   []string `json:"adTitle"`
	AdCopy    []string `json:"adCopy"`
	AdImage   []string `json:"adImage"`
	Objective string   `json:"objective"`
}

// TODO: implement
func (request newCampaignRequest) isValid() (valid bool) {
	valid = true
	return valid
}

// NewCampaign creates a Campaign from a NewCampaignRequest
func NewCampaign(c newCampaignRequest) (campaign model.Campaign, err error) {
	ID, err := uuid.NewRandom()
	campaign = model.Campaign{
		ID:        ID.String(),
		Template:  c.Template,
		AdTitle:   c.AdTitle,
		AdCopy:    c.AdCopy,
		Objective: c.Objective,
		Status:    model.Paused,
	}
	if c.AdImage != nil {
		campaign.AdImage = c.AdImage
	}
	return campaign, err
}

// Product represents the product object
type Product struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
}

// Products wraps the product array
type Products []Product
