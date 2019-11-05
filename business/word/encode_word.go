package word
import (
	"github.com/cisordeng/beego/xenon"
	"nature/business/account"
)

func EncodeWord(word *Word) xenon.Map {
	if word == nil {
		return nil
	}

	user := account.EncodeUser(word.User)

	mapWord := xenon.Map{
		"id": word.Id,
		"user": user,
		"content": word.Content,
		"created_at": word.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapWord
}


func EncodeManyWord(words []*Word) []xenon.Map {
	mapWords := make([]xenon.Map, 0)
	for _, word := range words {
		mapWords = append(mapWords, EncodeWord(word))
	}
	return mapWords
}
