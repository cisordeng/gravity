package article

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	m_article "mango/models/article"
	"mango/rest"
)

func GetArticleById(ArticleId int) (article *Article, err error, be rest.BusinessError) {
	model := m_article.Article{}
	err = orm.NewOrm().QueryTable("article_article").Filter("Id", ArticleId).One(&model)
	Article := InitArticleFromModel(&model)
	return Article, err, rest.BusinessError{
		ErrCode: "orm:article:not_exits",
		ErrMsg:  fmt.Sprintf("id为%d的article不存在", ArticleId),
	}
}

func GetArticles() (articles []*Article, err error, be rest.BusinessError) {
	var models []m_article.Article
	_, err = orm.NewOrm().QueryTable("article_article").All(&models)
	for _, model := range models {
		articles = append(articles, InitArticleFromModel(&model))
	}
	return articles, err, rest.BusinessError{
		ErrCode: "orm:articles:get_fail",
		ErrMsg:  "获取articles失败",
	}
}