package service

import (
	"github.com/DACNPM-Group5/leesin/pkg/model"
	"time"
)

type previewArticle struct {
	Title    string
	Image    string
	Content  string
	Author   string
	CreateAt time.Time
}

func fromArticle(article model.Article) previewArticle {
	image := ""
	p := ""
	ps := article.Body.Paragraph
	if len(p) > 0 {
		p = ps[0].Content
		imgs := ps[0].Image
		if len(imgs) > 0 {
			image = imgs[0]
		}
	}

	return previewArticle{
		Title:    article.Title,
		Image:    image,
		Content:  p,
		Author:   article.Author,
		CreateAt: article.CreatedAt,
	}
}
