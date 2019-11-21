package word

import (
	"nature/common/dolphin"
	"nature/common/leo"
)

func Fill(words []*Word) {
	if len(words) == 0 || words[0] == nil {
		return
	}

	leo.FillUser(words)
	dolphin.FillReplies(words)
}