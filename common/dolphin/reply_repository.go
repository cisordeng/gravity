package dolphin

import (
	"github.com/cisordeng/beego/xenon"
)

func GetReplies(filters xenon.Map, orderExprs ...string ) []*Reply {
	params := xenon.Map{
		"page": 1,
		"count_per_page": 9999,
		"filters": filters,
	}
	if len(orderExprs) > 0 {
		params["orders"] = orderExprs
	}
	resp := xenon.Get("dolphin", "reply.replies", params)
	replies := make([]*Reply, 0)
	replyInterfaces := resp["data"].(xenon.Map)["replies"].([]interface{})
	for _, reply := range replyInterfaces {
		replies = append(replies, InitReplyFromMap(reply.(xenon.Map)))
	}
	return replies
}
