package comment

import (
	"github.com/cisordeng/beego/xenon"
	"nature/business"
	"nature/business/account"
)

func Fill(resources []interface{}, restResource string) {
	resourceIds := make([]int, 0)
	for _, resource := range resources {
		resourceIds = append(resourceIds, business.GetValueByName(resource, "Id").(int))
	}

	comments := GetComments(xenon.Map{
		"resource_id__in": resourceIds,
		"resource_type": restResource,
	}, "-created_at")
	FillComment(comments)

	resourceId2comments := make(map[int][]*Comment)
	for _, comment := range comments {
		resourceId2comments[comment.ResourceId] = append(resourceId2comments[comment.ResourceId], comment)
	}

	for _, resource := range resources {
		if comments, ok := resourceId2comments[business.GetValueByName(resource, "Id").(int)]; ok {
			business.SetValueByName(resource, "Comments", comments)
		}
	}
	return
}

func FillComment(comments []*Comment) {
	if len(comments) == 0 || comments[0] == nil {
		return
	}

	fillUser(comments)
	fillComment(comments)
}

func fillUser(comments []*Comment) {
	userIds := make([]int, 0)
	for _, Comment := range comments {
		userIds = append(userIds, Comment.UserId)
	}

	users := account.GetUsers(xenon.Map{
		"id__in": userIds,
	})

	id2user := make(map[int]*account.User)
	for _, user := range users {
		id2user[user.Id] = user
	}

	for _, Comment := range comments {
		if user, ok := id2user[Comment.UserId]; ok {
			Comment.User = user
		}
	}
	return
}

func fillComment(comments []*Comment) {
	commentIds := make([]int, 0)
	for _, comment := range comments {
		commentIds = append(commentIds, comment.CommentId)
	}

	iComments := GetComments(xenon.Map{
		"id__in": commentIds,
	})

	id2comment := make(map[int]*Comment)
	for _, comment := range iComments {
		id2comment[comment.Id] = comment
	}

	for _, comment := range comments {
		if comment, ok := id2comment[comment.CommentId]; ok {
			comment.Comment = comment
		}
	}
	return
}
