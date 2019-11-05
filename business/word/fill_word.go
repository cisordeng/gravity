package word
import (
	"github.com/cisordeng/beego/xenon"

	"nature/business/account"
)

func Fill(words []*Word) {
	if len(words) == 0 || words[0] == nil {
		return
	}

	fillUser(words)
}


func fillUser(words []*Word) {
	userIds := make([]int, 0)
	for _, word := range words {
		userIds = append(userIds, word.UserId)
	}

	users := account.GetUsers(xenon.Map{
		"id__in": userIds,
	})

	id2user := make(map[int]*account.User)
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
