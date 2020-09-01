package utils

import (
	"github.com/tidwall/gjson"
)

func FindJsonKeys(result gjson.Result, lastKey string) []string {
	keys := []string{}

	if result.IsObject() {
		// collect keys
		for s := range result.Map() {
			keys = append(keys, formatedFirst(lastKey) + s)
			other := FindJsonKeys(result.Get(s), s)
			if len(other) > 0 {
				for i := range other {
					keys = append(keys, formatedFirst(lastKey) + other[i])
				}
			}
 		}
		return keys
	} else {
		return keys
	}
}

func formatedFirst(first string) string {
	if len(first) > 0 {
		return first + "."
	}
	return first
}