package rhythm

import (
	"github.com/cisordeng/beego/xenon"
	bRhythm "nature/business/rhythm"
)

type RhythmSets struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(RhythmSets))
}

func (this *RhythmSets) Resource() string {
	return "rhythm.rhythm_sets"
}

func (this *RhythmSets) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *RhythmSets) Get() {
	bCtx := this.GetBusinessContext()
	rhythmSets := bRhythm.GetRhythmSets(bCtx, xenon.Map{}, "-index")
	bRhythm.Fill(bCtx, rhythmSets, xenon.FillOption{
		"with_rhythm": true,
	})
	data := bRhythm.EncodeManyRhythmSet(rhythmSets)
	this.ReturnJSON(xenon.Map{
		"rhythm_sets": data,
	})
}
