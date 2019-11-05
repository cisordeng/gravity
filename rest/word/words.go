package word

import (
	"github.com/cisordeng/beego/xenon"

	bWord "nature/business/word"
)

type Words struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Words))
}

func (this *Words) Resource() string {
	return "word.words"
}

func (this *Words) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *Words) Get() {
	page := this.GetPage()
	words, pageInfo := bWord.GetPagedWords(page, xenon.Map{}, "-created_at")
	bWord.Fill(words)
	data := bWord.EncodeManyWord(words)
	this.ReturnJSON(xenon.Map{
		"words": data,
		"page_info": pageInfo.ToMap(),
	})
}
