package packs

type PackMcData struct {
	Pack struct {
		PackFormat  int    `json:"pack_format"`
		Description string `json:"description"`
	} `json:"pack"`
}
