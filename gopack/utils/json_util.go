package utils

import (
	"github.com/tidwall/gjson"
	"strconv"
)

var fileCache = make(map[string][]string)

func ResetFileCache() {
	fileCache = make(map[string][]string)
}

func FindJsonKeys(result gjson.Result, id string) []string {
	keys, found := fileCache[id]

	if found {
		return keys
	} else {
		keys = FindJsonKeysa(result, "")
		fileCache[id] = keys
		return keys
	}
}

func FindJsonKeysa(result gjson.Result, lastKey string) []string {
	keys := []string{}

	if result.IsObject() {
		// collect keys
		for s := range result.Map() {
			keys = append(keys, formatedFirst(lastKey) + s)
			other := FindJsonKeysa(result.Get(s), s)
			if len(other) > 0 {
				for i := range other {
					keys = append(keys, formatedFirst(lastKey) + other[i])
				}
			}
 		}
	} else if result.IsArray() {
		// collect keys
		a := result.Array()
		for i := range a {
			keys = append(keys, formatedFirst(lastKey) + strconv.Itoa(i))
			other := FindJsonKeysa(a[i], strconv.Itoa(i))
			if len(other) > 0 {
				for i2 := range other {
					keys = append(keys, formatedFirst(lastKey) + other[i2])
				}
			}
		}
	}
	return keys
}

func formatedFirst(first string) string {
	if len(first) > 0 {
		return first + "."
	}
	return first
}