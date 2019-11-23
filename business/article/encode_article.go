package article

import (
	"github.com/cisordeng/beego/xenon"

	"nature/common/dolphin"
	"nature/common/leo"
)

func EncodeArticle(article *Article) xenon.Map {
	if article == nil {
		return nil
	}

	user := leo.EncodeUser(article.User)
	rReplies := dolphin.EncodeManyReply(article.Replies)

	mapArticle := xenon.Map{
		"id": article.Id,
		"user": user,
		"replies": rReplies,
		"title": article.Title,
		"content": article.Content,
		"created_at": article.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapArticle
}


func EncodeManyArticle(articles []*Article) []xenon.Map {
	mapArticles := make([]xenon.Map, 0)
	for _, article := range articles {
		mapArticles = append(mapArticles, EncodeArticle(article))
	}
	return mapArticles
}
