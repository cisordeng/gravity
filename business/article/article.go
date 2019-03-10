package article

import (
	"fmt"
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"
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

func NewArticle(Title string, Content string) (article *Article, error xenon.Error) {
	model := m_article.Article{
		Title: Title,
		Content: Content,
	}
	_, error.Inner = orm.NewOrm().Insert(&model)
	error.Business = &xenon.BusinessError{
		ErrCode: "article:create_failed",
		ErrMsg:  fmt.Sprintf("创建%sarticle失败", Title),
	}
	return InitArticleFromModel(&model), error
}