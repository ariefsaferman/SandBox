package response

import (
	"url-shortener/internal/component/config"
	"url-shortener/internal/model"
)

type ItemGetOneBySlugResponse struct {
	Id       int64  `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func (r *ItemGetOneBySlugResponse) ConvertFromModel(cfg config.Config, m *model.Item) {
	r.Id = m.Id
	r.LongUrl = m.LongUrl
	r.Slug = m.Slug
	r.Title = m.Title
	r.Content = m.Content
}

type ItemCreateOneResponse struct {
	Id       int64  `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	Slug     string `json:"slug"`
	Title    string `form:"title"`
	Content  string `form:"content"`
}

func (r *ItemCreateOneResponse) ConvertFromModel(m *model.Item) {
	r.Id = m.Id
	r.LongUrl = m.LongUrl
	r.Slug = m.Slug
	r.Title = m.Title
	r.Content = m.Content
}
