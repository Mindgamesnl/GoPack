package gopack

type PackMcMeta struct {
	Pack struct {
		PackFormat  int    `json:"pack_format"`
		Description string `json:"description"`
	} `json:"pack"`
}
