package comment

import (
	"github.com/cisordeng/beego/xenon"

	bComment "nature/business/comment"
)

type Comments struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Comments))
}

func (this *Comments) Resource() string {
	return "comment.comments"
}

func (this *Comments) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *Comments) Get() {
	page := this.GetPage()
	comments, pageInfo := bComment.GetPagedComments(page, xenon.Map{}, "-created_at")
	data := bComment.EncodeManyComment(comments)
	this.ReturnJSON(xenon.Map{
		"comments": data,
		"page_info": pageInfo.ToMap(),
	})
}

