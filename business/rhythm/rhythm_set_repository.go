package rhythm

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mRhythm "nature/model/rhythm"
)

func GetOneRhythmSet(filters xenon.Map) *RhythmSet {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.RhythmSet{})

	var model mRhythm.RhythmSet
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}

	err := qs.One(&model)
	xenon.PanicNotNilError(err, "raise:rhythm_set:not_exits", "歌单不存在")
	return InitRhythmSetFromModel(&model)
}

func GetRhythmSets(filters xenon.Map, orderExprs ...string ) []*RhythmSet {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.RhythmSet{})

	var models []*mRhythm.RhythmSet
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	_, err := qs.All(&models)
	xenon.PanicNotNilError(err)


	rhythmSets := make([]*RhythmSet, 0)
	for _, model := range models {
		rhythmSets = append(rhythmSets, InitRhythmSetFromModel(model))
	}
	return rhythmSets
}

func GetRhythmSet(id int) *RhythmSet {
	return GetOneRhythmSet(xenon.Map{
		"id": id,
	})
}
