package word

import (
	"github.com/cisordeng/beego/xenon"

	"nature/common/leo"
	bWord "nature/business/word"
)

type Word struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Word))
}

func (this *Word) Resource() string {
	return "word.word"
}

func (this *Word) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"id",
		},
		"PUT":  []string{
			"token",
			"content",
		},
		"POST": []string{
			"token",
			"id",
			"content",
		},
	}
}

func (this *Word) Get() {
	id, _ := this.GetInt("id", 0)

	word := bWord.GetWordById(id)
	bWord.Fill([]*bWord.Word{ word })
	data := bWord.EncodeWord(word)
	this.ReturnJSON(data)
}

func (this *Word) Put() {
	content := this.GetString("content")

	user := leo.User{}
	this.GetUserFromToken(&user)

	word := bWord.NewWord(user, content)
	bWord.Fill([]*bWord.Word{ word })
	data := bWord.EncodeWord(word)
	this.ReturnJSON(data)
}

func (this *Word) Post() {
	id, _ := this.GetInt("id", 0)
	content := this.GetString("content")

	word := bWord.GetWordById(id)
	word.Update(content)
	bWord.Fill([]*bWord.Word{ word })
	data := bWord.EncodeWord(word)
	this.ReturnJSON(data)
}