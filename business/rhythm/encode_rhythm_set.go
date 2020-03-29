package rhythm

import (
	"github.com/cisordeng/beego/xenon"
)

func EncodeRhythmSet(rhythmSet *RhythmSet) xenon.Map {
	if rhythmSet == nil {
		return nil
	}

	mapRhythmSet := xenon.Map{
		"id": rhythmSet.Id,
		"nid": rhythmSet.Nid,

		"name": rhythmSet.Name,
		"avatar": rhythmSet.Avatar,
		"played_count": rhythmSet.PlayedCount,
		"created_at": rhythmSet.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapRhythmSet
}


func EncodeManyRhythmSet(rhythmSets []*RhythmSet) []xenon.Map {
	mapRhythmSets := make([]xenon.Map, 0)
	for _, rhythmSet := range rhythmSets {
		mapRhythmSets = append(mapRhythmSets, EncodeRhythmSet(rhythmSet))
	}
	return mapRhythmSets
}

