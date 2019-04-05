package rhythm

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"

	mRhythm "gravity/model/rhythm"
)

func GetRhythmSets(ctx *xenon.Ctx, filters xenon.Map, orderExprs ...string ) []*RhythmSet {
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
	xenon.RaiseError(ctx, err)


	rhythmSets := make([]*RhythmSet, 0)
	for _, model := range models {
		rhythmSets = append(rhythmSets, InitRhythmSetFromModel(model))
	}
	return rhythmSets
}

func GetRhythmSet(ctx *xenon.Ctx, id int) *RhythmSet {
	rhythmSets := GetRhythmSets(ctx, xenon.Map{
		"id": id,
	}, "-id")
	if len(rhythmSets) > 0 {
		return rhythmSets[0]
	} else {
		return nil
	}
}
