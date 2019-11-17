package business

import "reflect"

func init() {
}

func SetValueByName(resource interface{}, name string, value interface{}) {
	v := reflect.ValueOf(resource).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Name == name {
			v.FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}
}

func GetValueByName(resource interface{}, name string) interface{} {
	v := reflect.ValueOf(resource).Elem()
	return v.FieldByName(name).Interface()
}