package rhythm

import (
	"github.com/cisordeng/beego/xenon"
	bRhythm "nature/business/rhythm"
)

type Rhythms struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Rhythms))
}

func (this *Rhythms) Resource() string {
	return "rhythm.rhythms"
}

func (this *Rhythms) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"rhythm_set_id",
		},
	}
}

func (this *Rhythms) Get() {
	rhythmSetId, _ := this.GetInt("rhythm_set_id", 0)

	page := this.GetPage()
	rhythmSet := bRhythm.GetRhythmSet(rhythmSetId)
	rhythms, pageInfo := bRhythm.GetPagedRhythmsByRhythmSet(rhythmSet, page)
	data := bRhythm.EncodeManyRhythm(rhythms)
	this.ReturnJSON(xenon.Map{
		"rhythms": data,
		"page_info": pageInfo.ToMap(),
	})
}
