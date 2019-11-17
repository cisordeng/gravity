package comment

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mComment "nature/model/comment"
)

func GetOneComment(filters xenon.Map) *Comment {
	o := orm.NewOrm()
	qs := o.QueryTable(&mComment.Comment{})

	var model mComment.Comment
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	if err != nil {
		return nil
	}
	//xenon.PanicNotNilError(err, "raise:comment:not_exits", "comment不存在")
	return InitCommentFromModel(&model)
}

func GetComments(filters xenon.Map, orderExprs ...string ) []*Comment {
	o := orm.NewOrm()
	qs := o.QueryTable(&mComment.Comment{})

	var models []*mComment.Comment
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	comments := make([]*Comment, 0)
	for _, model := range models {
		comments = append(comments, InitCommentFromModel(model))
	}
	return comments
}

func GetPagedComments(page *xenon.Paginator, filters xenon.Map, orderExprs ...string ) ([]*Comment, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mComment.Comment{})

	var models []*mComment.Comment
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)

	comment := make([]*Comment, 0)
	for _, model := range models {
		comment = append(comment, InitCommentFromModel(model))
	}
	return comment, pageInfo
}

func GetCommentById(id int) *Comment {
	return GetOneComment(xenon.Map{
		"id": id,
	})
}
