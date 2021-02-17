package search

import (
	"reflect"
)

func toMap(sr interface{}) map[string]interface{} {
	v := reflect.ValueOf(sr)
	values := make(map[string]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		// if field is exported
		if v.Field(i).CanInterface() {
			// assign map's key as json tag, and value as is
			values[v.Type().Field(i).Tag.Get("json")] = v.Field(i).Interface()
		}
	}
	return values
}
