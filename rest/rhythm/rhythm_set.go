package rhythm

import (
	"github.com/cisordeng/beego/xenon"
	bRhythm "nature/business/rhythm"
)

type RhythmSet struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(RhythmSet))
}

func (this *RhythmSet) Resource() string {
	return "rhythm.rhythm_set"
}

func (this *RhythmSet) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"id",
		},
	}
}

func (this *RhythmSet) Get() {
	id, _ := this.GetInt("id", 0)

	bCtx := this.GetBusinessContext()
	rhythmSet := bRhythm.GetRhythmSet(bCtx, id)
	bRhythm.Fill(bCtx, []*bRhythm.RhythmSet{ rhythmSet }, xenon.FillOption{
		"with_rhythm": true,
	})
	data := bRhythm.EncodeRhythmSet(rhythmSet)
	this.ReturnJSON(xenon.Map{
		"rhythm_set": data,
	})
}
