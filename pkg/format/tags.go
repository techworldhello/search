package format

import "reflect"

func getHeaderTags(entity interface{}) (tags []string) {
	v := reflect.ValueOf(entity)

	for i := 0; i < v.NumField(); i++ {
		header := v.Type().Field(i).Tag.Get("header")
		if header != "" {
			tags = append(tags, header)
		}
	}
	return tags
}
