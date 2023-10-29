package models

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionData struct {
	Data []Option `json:"data"`
}
