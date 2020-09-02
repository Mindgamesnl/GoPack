package gopack

import "github.com/tidwall/gjson"

var typeCache = make(map[string]gjson.Type)

// getting types is slow as shit
// so caching it helps a lot
func getTypeCache(key string, res gjson.Result) gjson.Type {
	fc, found := typeCache[key]
	if found {
		return fc
	}
	ft := res.Type
	typeCache[key] = ft
	return ft
}
