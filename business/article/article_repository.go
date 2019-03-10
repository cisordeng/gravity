package article

import (
	"fmt"
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"
	m_article "mango/models/article"
)

func GetArticleById(ArticleId int) (article *Article, error xenon.Error) {
	model := m_article.Article{}
	error.Inner = orm.NewOrm().QueryTable("article_article").Filter("Id", ArticleId).One(&model)
	error.Business = &xenon.BusinessError{
		"orm:article:not_exits",
		fmt.Sprintf("id为%d的article不存在", ArticleId),
	}
	article = InitArticleFromModel(&model)
	return article, error
}

func GetArticles() (articles []*Article, error xenon.Error) {
	var models []m_article.Article
	_, error.Inner = orm.NewOrm().QueryTable("article_article").All(&models)
	error.Business = &xenon.BusinessError{
		ErrCode: "orm:articles:get_fail",
		ErrMsg:  "获取articles失败",
	}
	for _, model := range models {
		articles = append(articles, InitArticleFromModel(&model))
	}
	return articles, error
}