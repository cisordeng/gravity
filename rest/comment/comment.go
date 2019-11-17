package comment

import (
	"github.com/cisordeng/beego/xenon"

	bUser "nature/business/account"
	bComment "nature/business/comment"
	bWord "nature/business/word"
)

type Comment struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Comment))
}

func (this *Comment) Resource() string {
	return "comment.comment"
}

func (this *Comment) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{
			"id",
		},
		"PUT":  []string{
			"token",
			"resource_type",
			"resource_id",
			"comment_id",
			"content",
		},
	}
}

func (this *Comment) Get() {
	id, _ := this.GetInt("id", 0)

	comment := bComment.GetCommentById(id)
	data := bComment.EncodeComment(comment)
	this.ReturnJSON(data)
}

func (this *Comment) Put() {
	resourceType := this.GetString("resource_type")
	resourceId, _ := this.GetInt("resource_id", 0)
	commentId, _ := this.GetInt("comment_id", 0)
	content := this.GetString("content")

	user := bUser.User{}
	this.GetUserFromToken(&user)

	comment := bComment.GetCommentById(commentId)
	word := bWord.GetWordById(resourceId)

	iComment := bComment.NewComment(user, resourceType, word, comment, content)
	bComment.FillComment([]*bComment.Comment{ iComment })
	data := bComment.EncodeComment(iComment)
	this.ReturnJSON(data)
}
