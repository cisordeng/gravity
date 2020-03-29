package rhythm

import (
	"github.com/cisordeng/beego/xenon"
	bRhythm "nature/business/rhythm"
)

type Rhythm struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Rhythm))
}

func (this *Rhythm) Resource() string {
	return "rhythm.rhythm"
}

func (this *Rhythm) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"id",
		},
	}
}

func (this *Rhythm) Get() {
	id, _ := this.GetInt("id", 0)
	rhythm := bRhythm.GetRhythm(id)
	data := bRhythm.EncodeRhythm(rhythm)
	this.ReturnJSON(data)
}
