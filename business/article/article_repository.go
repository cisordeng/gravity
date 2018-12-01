package article

import (
	"github.com/astaxie/beego/orm"
	m_article "mango/models/article"
)

func GetArticleById(ArticleId int) (article *Article) {
	model := m_article.Article{}
	orm.NewOrm().QueryTable("article_article").Filter("Id", ArticleId).One(&model)
	Article := InitArticleFromModel(&model)
	return Article
}