package comment
import (
	"github.com/cisordeng/beego/xenon"
	"nature/business/account"
)

func EncodeComment(comment *Comment) xenon.Map {
	if comment == nil {
		return nil
	}

	rUser := account.EncodeUser(comment.User)
	rComment := EncodeComment(comment.Comment)

	mapComment := xenon.Map{
		"id": comment.Id,
		"user": rUser,
		"comment": rComment,
		"content": comment.Content,
		"created_at": comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return mapComment
}


func EncodeManyComment(comments []*Comment) []xenon.Map {
	mapComments := make([]xenon.Map, 0)
	for _, comment := range comments {
		mapComments = append(mapComments, EncodeComment(comment))
	}
	return mapComments
}
