package rhythm

import (
	"github.com/cisordeng/beego/orm"
	"github.com/cisordeng/beego/xenon"
	"sort"

	mRhythm "gravity/model/rhythm"
)

func Fill(ctx *xenon.Ctx, rhythmSets []*RhythmSet, option xenon.FillOption) {
	ids := make([]int, 0)
	for _, rhythmSet := range rhythmSets {
		ids = append(ids, rhythmSet.Id)
	}
	if enableOption, ok := option["with_rhythm"]; ok && enableOption {
		fillRhythm(ctx, rhythmSets, ids)
	}
}

func fillRhythm(ctx *xenon.Ctx, rhythmSets []*RhythmSet, ids []int) {
	o := orm.NewOrm()
	qs := o.QueryTable(&mRhythm.RhythmSetRhythm{})
	qs.Filter(xenon.Map{
		"rhythm_set_id__in": ids,
	})

	var models []*mRhythm.RhythmSetRhythm
	_, err := qs.All(&models)
	xenon.RaiseError(ctx, err)

	rhythmIds := make([]int, 0)
	id2rhythmIds := make(map[int][]int, 0)
	for _, model := range models {
		rhythmIds = append(rhythmIds, model.RhythmId)
		id2rhythmIds[model.RhythmSetId] = append(id2rhythmIds[model.RhythmSetId], model.RhythmId)
	}

	rhythmId2rhythm := make(map[int]*Rhythm, 0)
	rhythms := GetRhythms(ctx, xenon.Map{
		"id__in": rhythmIds,
	})
	for _, rhythm := range rhythms {
		rhythmId2rhythm[rhythm.Id] = rhythm
	}

	for _, rhythmSet := range rhythmSets {
		rhythms = make([]*Rhythm, 0)
		for _, rhythmId := range id2rhythmIds[rhythmSet.Id] {
			rhythms = append(rhythms, rhythmId2rhythm[rhythmId])
		}
		sort.Slice(rhythms, func(i int, j int) bool {
			return rhythms[i].Id < rhythms[j].Id
		})
		rhythmSet.Rhythms = rhythms
	}
}

