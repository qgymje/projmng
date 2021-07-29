package service

import "projmng/pkg/entity"

type ContentTemplateService struct {
	contentTemplate entity.IContentTemplate
}

func (cts *ContentTemplateService) Create(ct *entity.ContentTemplate) error {
	return cts.contentTemplate.Create(ct)
}

func (cts *ContentTemplateService) Get(ct entity.ContentTemplate) (*entity.ContentTemplate, error) {
	panic("")
}
