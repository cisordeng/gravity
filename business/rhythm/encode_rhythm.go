package rhythm

import (
	"net/http"
	"strings"

	"github.com/cisordeng/beego/xenon"
)

func EncodeRhythm(rhythm *Rhythm) xenon.Map {

	url := getRedirectUrl(rhythm.Url)

	mapRhythm := xenon.Map{
		"id": rhythm.Id,
		"nid": rhythm.Nid,

		"name": rhythm.Name,
		"avatar": rhythm.Avatar,
		"url": strings.Replace(url, "http://", "https://", 1),
		"lyric": rhythm.Lyric,
		"translated_lyric": rhythm.TranslatedLyric,
		"singer_name": rhythm.SingerName,
		"played_count": rhythm.PlayedCount,
		"created_at": rhythm.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapRhythm
}


func EncodeManyRhythm(rhythms []*Rhythm) []xenon.Map {
	mapRhythms := make([]xenon.Map, 0)
	for _, rhythm := range rhythms {
		mapRhythms = append(mapRhythms, EncodeRhythm(rhythm))
	}
	return mapRhythms
}

func getRedirectUrl(url string) string {
	resp, err := http.Head(url)
	xenon.PanicNotNilError(err)
	return resp.Request.URL.String()
}
