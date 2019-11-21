package dolphin

import (
	"reflect"

	"github.com/cisordeng/beego/xenon"
)

func FillReplies(resources interface{}) {
	resourceIds := make([]int, 0)
	for i := 0; i < reflect.ValueOf(resources).Len(); i ++ {
		resource := reflect.ValueOf(resources).Index(i)
		resourceIds = append(resourceIds, resource.Elem().FieldByName("Id").Interface().(int))
	}

	bytes := reflect.TypeOf(resources).String()[3:]
	resourceType := ""
	for i, b := range bytes {
		if b >= 'A' && b <= 'Z' {
			if i - 1 >= 0 && bytes[i - 1] != '.' {
				resourceType += "_"
			}
			resourceType += string(b + 32)
		} else {
			resourceType += string(b)
		}
	}
	replies := GetReplies(xenon.Map{
		"resource_id__in": resourceIds,
		"resource_type": resourceType,
	})


	resourceId2replies := make(map[int][]*Reply)
	for _, reply := range replies {
		resourceId2replies[reply.ResourceId] = append(resourceId2replies[reply.ResourceId], reply)
	}

	for i := 0; i < reflect.ValueOf(resources).Len(); i ++ {
		resource := reflect.ValueOf(resources).Index(i)
		if replies, ok := resourceId2replies[resource.Elem().FieldByName("Id").Interface().(int)]; ok {
			resource.Elem().FieldByName("Replies").Set(reflect.ValueOf(replies))
		}
	}
	return
}
