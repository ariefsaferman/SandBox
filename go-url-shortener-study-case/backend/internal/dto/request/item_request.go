package request

import (
	"url-shortener/internal/model"
)

type ItemCreateOneRequest struct {
	LongUrl string `json:"long_url"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *ItemCreateOneRequest) ConvertToModel() model.Item {
	return model.Item{
		LongUrl: p.LongUrl,
		Slug:    p.Slug,
		Title:   p.Title,
		Content: p.Content,
	}
}
