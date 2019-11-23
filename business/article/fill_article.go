package article

import (
	"nature/common/dolphin"
	"nature/common/leo"
)

func Fill(articles []*Article) {
	if len(articles) == 0 || articles[0] == nil {
		return
	}

	leo.FillUser(articles)
	dolphin.FillReplies(articles)
}