package rhythm

import (
	"time"

	mRhythm "gravity/model/rhythm"
)

type Rhythm struct {
	Id int // id
	Nid int64

	Name string
	Avatar string
	Url string
	Lyric string
	TranslatedLyric string
	SingerName string
	PlayedCount int
	CreatedAt time.Time
}

func init() {
}

func InitRhythmFromModel(model *mRhythm.Rhythm) *Rhythm {
	instance := new(Rhythm)
	instance.Id = model.Id
	instance.Nid = model.Nid
	instance.Name = model.Name
	instance.Avatar = model.Avatar
	instance.Url = model.Url
	instance.Lyric = model.Lyric
	instance.TranslatedLyric = model.TranslatedLyric
	instance.SingerName = model.SingerName
	instance.PlayedCount = model.PlayedCount
	instance.CreatedAt = model.CreatedAt

	return instance
}