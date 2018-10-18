package handler

import (
	"errors"
	"strings"

	"github.com/danew/ad-campaign/model"
)

type inMemoryRepository struct {
	campaigns []model.Campaign
	templates []model.Template
}

// NewRepository creates a new in-memory repository
func newInMemoryRepository() *inMemoryRepository {
	repo := &inMemoryRepository{}
	repo.campaigns = []model.Campaign{}
	repo.templates = []model.Template{}
	return repo
}

func (repo *inMemoryRepository) addCampaign(campaign model.Campaign) (err error) {
	repo.campaigns = append(repo.campaigns, campaign)
	return err
}

func (repo *inMemoryRepository) getCampaigns() (campaigns []model.Campaign, err error) {
	campaigns = repo.campaigns
	return
}

func (repo *inMemoryRepository) getPublishedCampaigns() (campaigns []model.Campaign, err error) {
	items := make([]model.Campaign, 0)
	for _, campaign := range repo.campaigns {
		if strings.Compare(campaign.Status, model.Implemented) == 0 {
			items = append(items, campaign)
		}
	}
	return items, err
}

func (repo *inMemoryRepository) getCampaign(id string) (campaign model.Campaign, err error) {
	found := false
	for _, target := range repo.campaigns {
		if strings.Compare(target.ID, id) == 0 {
			campaign = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find campaign in repository")
	}
	return campaign, err
}

func (repo *inMemoryRepository) updateCampaign(id string, campaign model.Campaign) (err error) {
	found := false
	for k, v := range repo.campaigns {
		if strings.Compare(v.ID, id) == 0 {
			repo.campaigns[k] = campaign
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find campaign in repository")
	}
	return
}

func (repo *inMemoryRepository) addTemplate(template model.Template) (err error) {
	repo.templates = append(repo.templates, template)
	return err
}

func (repo *inMemoryRepository) getTemplates() (templates []model.Template, err error) {
	templates = repo.templates
	return
}

func (repo *inMemoryRepository) getTemplate(id string) (template model.Template, err error) {
	found := false
	for _, target := range repo.templates {
		if strings.Compare(target.ID, id) == 0 {
			template = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find template in repository")
	}
	return template, err
}
