package article

import (
	"github.com/cisordeng/beego/xenon"
)

func EncodeArticle(article *Article) xenon.Map {
	mapArticle := xenon.Map{
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