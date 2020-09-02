package sources

import "encoding/json"

func Get14Blocks() map[string]string {
	a := make(map[string]string)
	_ = json.Unmarshal([]byte(blocks14), &a)
	return a
}

func GetBlocks() map[string]string {
	a := make(map[string]string)
	_ = json.Unmarshal([]byte(blocks), &a)
	return a
}

func GetItems() map[string]string {
	a := make(map[string]string)
	_ = json.Unmarshal([]byte(items), &a)
	return a
}

func Get14Items() map[string]string {
	a := make(map[string]string)
	_ = json.Unmarshal([]byte(items14), &a)
	return a
}

func GetLang() map[string]string {
	a := make(map[string]string)
	_ = json.Unmarshal([]byte(lang), &a)
	return a
}

func GetLang14() map[string]string {
	a := make(map[string]string)
	_ = json.Unmarshal([]byte(lang14), &a)
	return a
}
