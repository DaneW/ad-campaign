package handler

import (
	"github.com/danew/service/model"
)

type inMemoryRepository struct {
	templates []model.Template
}

// NewRepository creates a new in-memory repository
func newInMemoryRepository() *inMemoryRepository {
	repo := &inMemoryRepository{}
	repo.templates = []model.Template{}
	return repo
}

func (repo *inMemoryRepository) getTemplates() (templates []model.Template, err error) {
	templates = repo.templates
	return
}
