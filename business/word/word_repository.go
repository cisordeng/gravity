package word

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mWord "nature/model/word"
)

func GetOneWord(filters xenon.Map) *Word {
	o := orm.NewOrm()
	qs := o.QueryTable(&mWord.Word{})

	var model mWord.Word
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	xenon.PanicNotNilError(err, "raise:word:not_exits", "word不存在")
	return InitWordFromModel(&model)
}

func GetWords(filters xenon.Map, orderExprs ...string ) []*Word {
	o := orm.NewOrm()
	qs := o.QueryTable(&mWord.Word{})

	var models []*mWord.Word
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	words := make([]*Word, 0)
	for _, model := range models {
		words = append(words, InitWordFromModel(model))
	}
	return words
}

func GetPagedWords(page *xenon.Paginator, filters xenon.Map, orderExprs ...string ) ([]*Word, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mWord.Word{})

	var models []*mWord.Word
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)

	words := make([]*Word, 0)
	for _, model := range models {
		words = append(words, InitWordFromModel(model))
	}
	return words, pageInfo
}

func GetWordById(id int) *Word {
	return GetOneWord(xenon.Map{
		"id": id,
	})
}
