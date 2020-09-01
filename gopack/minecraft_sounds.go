package gopack

import "encoding/json"

type Sound struct {
	Sounds []struct {
		Name   string `json:"name"`
		Stream bool   `json:"stream"`
	} `json:"sounds"`
	Category string `json:"category"`
}

type Sounds struct {
	Sounds map[string]Sound
}

func SoundsFromJson(input string) Sounds {
	m := make(map[string]Sound)
	_ = json.Unmarshal([]byte(input), &m)
	return Sounds{
		Sounds: m,
	}
}
