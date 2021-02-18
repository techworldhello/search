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

func getJsonTagsInMap(entities ...interface{}) (tags map[string][]string) {
	for _, entity := range entities {
		v := reflect.ValueOf(entity)
		name := reflect.TypeOf(entity).Name()

		for i := 0; i < v.NumField(); i++ {
			if tags == nil {
				tags = make(map[string][]string)
			}
			tags[name] = append(tags[name], v.Type().Field(i).Tag.Get("json"))
		}
	}

	return tags
}
