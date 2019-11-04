package article

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mArticle "nature/model/article"
)

func GetOneArticle(filters xenon.Map) *Article {
	o := orm.NewOrm()
	qs := o.QueryTable(&mArticle.Article{})

	var model mArticle.Article
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	xenon.PanicNotNilError(err, "raise:article:not_exits", "article不存在")
	return InitArticleFromModel(&model)
}

func GetArticles(filters xenon.Map, orderExprs ...string ) []*Article {
	o := orm.NewOrm()
	qs := o.QueryTable(&mArticle.Article{})

	var models []*mArticle.Article
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	articles := make([]*Article, 0)
	for _, model := range models {
		articles = append(articles, InitArticleFromModel(model))
	}
	return articles
}

func GetPagedArticles(page *xenon.Paginator, filters xenon.Map, orderExprs ...string ) ([]*Article, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mArticle.Article{})

	var models []*mArticle.Article
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)

	articles := make([]*Article, 0)
	for _, model := range models {
		articles = append(articles, InitArticleFromModel(model))
	}
	return articles, pageInfo
}

func GetArticleById(id int) *Article {
	return GetOneArticle(xenon.Map{
		"id": id,
	})
}
