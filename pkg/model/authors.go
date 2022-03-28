package model

//Authors model
type Authors struct {
	Authors []Author `json:"Authors"`
}

//Author detail model
type Author struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
}
