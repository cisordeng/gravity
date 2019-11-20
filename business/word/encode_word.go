package word

import (
	"github.com/cisordeng/beego/xenon"
	"nature/common/dolphin"

	"nature/common/leo"
)

func EncodeWord(word *Word) xenon.Map {
	if word == nil {
		return nil
	}

	user := leo.EncodeUser(word.User)
	rReplies := dolphin.EncodeManyReply(word.Replies)

	mapWord := xenon.Map{
		"id": word.Id,
		"user": user,
		"replies": rReplies,
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
