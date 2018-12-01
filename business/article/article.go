package article

import (
	"github.com/astaxie/beego/orm"
	m_article "mango/models/article"
	"time"
)

type Article struct {
	Id int
	Title string
	Content string
	CreatedAt time.Time
}

func init() {
}

func InitArticleFromModel(model *m_article.Article) *Article {
	instance := new(Article)
	instance.Id = model.Id
	instance.Title = model.Title
	instance.Content = model.Content
	instance.CreatedAt = model.CreatedAt

	return instance
}

func Create(Title string, Content string) *Article {
	model := m_article.Article{
		Title: Title,
		Content: Content,
	}
	orm.NewOrm().Insert(&model)
	return InitArticleFromModel(&model)
}