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
	return InitRhythmFromModel(&model)
}

func GetPagedRhythms(page *xenon.Paginator, filters xenon.Map, orderExprs ...string ) ([]*Rhythm, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.Rhythm{})

	var models []*mRhythm.Rhythm
	if len(filters) > 0 {
		qs = qs.Filter(filters)
	}
	if len(orderExprs) > 0 {
		qs = qs.OrderBy(orderExprs...)
	}

	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)

	rhythms := make([]*Rhythm, 0)
	for _, model := range models {
		rhythms = append(rhythms, InitRhythmFromModel(model))
	}
	return rhythms, pageInfo
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

func GetPagedRhythmsByRhythmSet(rhythmSet *RhythmSet, page *xenon.Paginator) ([]*Rhythm, xenon.PageInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.RhythmSetRhythm{})
	qs = qs.Filter(xenon.Map{
		"rhythm_set_id": rhythmSet.Id,
	})

	qs = qs.OrderBy("-index")

	var models []*mRhythm.RhythmSetRhythm
	pageInfo, err := xenon.Paginate(qs, page, &models)
	xenon.PanicNotNilError(err)
	rhythmIds := make([]int, 0)
	for _, model := range models {
		rhythmIds = append(rhythmIds, model.RhythmId)
	}
	iRhythms := GetRhythms(xenon.Map{
		"id__in": rhythmIds,
	})

	id2rhythm := make(map[int]*Rhythm, 0)
	for _, rhythm := range iRhythms {
		id2rhythm[rhythm.Id] = rhythm
	}

	rhythms := make([]*Rhythm, 0)
	for _, model := range models {
		rhythms = append(rhythms, id2rhythm[model.RhythmId])
	}
	return rhythms, pageInfo
}

func GetRhythm(id int) *Rhythm {
	return GetOneRhythm(xenon.Map{
		"id": id,
	})
}
