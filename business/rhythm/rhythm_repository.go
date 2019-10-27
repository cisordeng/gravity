package rhythm

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mRhythm "nature/model/rhythm"
)

func GetOneRhythm(filters xenon.Map) *Rhythm {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.Rhythm{})

	var model mRhythm.Rhythm
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	xenon.PanicNotNilError(err, "raise:rhythm:not_exits", "歌曲不存在")
	return nil
}

func GetRhythms(filters xenon.Map, orderExprs ...string ) []*Rhythm {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.Rhythm{})

	var models []*mRhythm.Rhythm
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	rhythms := make([]*Rhythm, 0)
	for _, model := range models {
		rhythms = append(rhythms, InitRhythmFromModel(model))
	}
	return rhythms
}

func GetRhythm(id int) *Rhythm {
	return GetOneRhythm(xenon.Map{
		"id": id,
	})
}
