package article

import (
	"mango/business"
)

func EncodeArticle(article *Article) business.Map {
	mapArticle := business.Map{
		"id": article.Id,
		"title": article.Title,
		"content": article.Content,
		"created_at": article.CreatedAt,
	}
	if article.Id !=0 {
		return mapArticle
	} else {
		return nil
	}
}