package article

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m_article "mango/models/article"
	"mango/rest"
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

func Create(Title string, Content string) (article *Article, err error, businessError rest.BusinessError) {
	model := m_article.Article{
		Title: Title,
		Content: Content,
	}
	Id, err := orm.NewOrm().Insert(&model)
	beego.NotNil(Id)
	return InitArticleFromModel(&model), err, rest.BusinessError{
		ErrCode: "article:create_failed",
		ErrMsg:  fmt.Sprintf("创建%sarticle失败", Title),
	}
}