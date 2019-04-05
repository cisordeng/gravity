package rhythm

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mRhythm "gravity/model/rhythm"
)

func GetRhythms(ctx *xenon.Ctx, filters xenon.Map, orderExprs ...string ) []*Rhythm {
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
	xenon.RaiseError(ctx, err)


	rhythms := make([]*Rhythm, 0)
	for _, model := range models {
		rhythms = append(rhythms, InitRhythmFromModel(model))
	}
	return rhythms
}

func GetRhythm(ctx *xenon.Ctx, id int) *Rhythm {
	rhythms := GetRhythms(ctx, xenon.Map{
		"id": id,
	}, "-id")
	if len(rhythms) > 0 {
		return rhythms[0]
	} else {
		return nil
	}
}
