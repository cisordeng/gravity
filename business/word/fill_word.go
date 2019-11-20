package word

import (
	"github.com/cisordeng/beego/xenon"

	"nature/common/dolphin"
	"nature/common/leo"
)

func Fill(words []*Word) {
	if len(words) == 0 || words[0] == nil {
		return
	}

	fillUser(words)
	fillReplies(words)
}


func fillUser(words []*Word) {
	userIds := make([]int, 0)
	for _, word := range words {
		userIds = append(userIds, word.UserId)
	}

	users := leo.GetUsers(xenon.Map{
		"id__in": userIds,
	})

	id2user := make(map[int]*leo.User)
	for _, user := range users {
		id2user[user.Id] = user
	}

	for _, word := range words {
		if user, ok := id2user[word.UserId]; ok {
			word.User = user
		}
	}
	return
}

func fillReplies(words []*Word) {
	wordIds := make([]int, 0)
	for _, word := range words {
		wordIds = append(wordIds, word.Id)
	}

	replies := dolphin.GetReplies(xenon.Map{
		"resource_id__in": wordIds,
		"resource_type": "word.word",
	})


	resourceId2replies := make(map[int][]*dolphin.Reply)
	for _, reply := range replies {
		resourceId2replies[reply.ResourceId] = append(resourceId2replies[reply.ResourceId], reply)
	}

	for _, word := range words {
		if replies, ok := resourceId2replies[word.Id]; ok {
			word.Replies = replies
		}
	}
	return
}
