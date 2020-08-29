package _12

type Face12 struct {
	Down struct {
		Uv      []int  `json:"uv"`
		Texture string `json:"texture"`
	} `json:"down"`
	Up struct {
		Uv      []int  `json:"uv"`
		Texture string `json:"texture"`
	} `json:"up"`
	North struct {
		Uv      []float64 `json:"uv"`
		Texture string    `json:"texture"`
	} `json:"north"`
	South struct {
		Uv      []float64 `json:"uv"`
		Texture string    `json:"texture"`
	} `json:"south"`
	West struct {
		Uv      []int  `json:"uv"`
		Texture string `json:"texture"`
	} `json:"west"`
	East struct {
		Uv      []int  `json:"uv"`
		Texture string `json:"texture"`
	} `json:"east"`
}

type Displays12 struct {
	Rotation    []int     `json:"rotation"`
	Translation []float64 `json:"translation"`
	Scale       []float64 `json:"scale"`
}

type BlockJson12 struct {
	Parent   string `json:"parent"`
	Comment  string `json:"__comment"`
	Textures struct {
		Particle string `json:"particle"`
		Texture  string `json:"texture"`
	} `json:"textures"`
	Elements []struct {
		From  []float64         `json:"from"`
		To    []float64         `json:"to"`
		Faces map[string]Face12 `json:"faces,omitempty"`
	} `json:"elements"`
	Display map[string]Displays12 `json:"display"`
}
