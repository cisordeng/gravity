package rhythm

import (
	"time"

	mRhythm "gravity/model/rhythm"
)

type RhythmSet struct {
	Id int
	Nid int64

	Name string
	Avatar string
	PlayedCount int
	CreatedAt time.Time

	Rhythms []*Rhythm
}

func init() {
}

func InitRhythmSetFromModel(model *mRhythm.RhythmSet) *RhythmSet {
	instance := new(RhythmSet)
	instance.Id = model.Id
	instance.Nid = model.Nid
	instance.Name = model.Name
	instance.Avatar = model.Avatar
	instance.PlayedCount = model.PlayedCount
	instance.CreatedAt = model.CreatedAt

	return instance
}