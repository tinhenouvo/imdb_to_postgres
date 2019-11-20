package models

// TitleAkas movie basic info
type TitleAkas struct {
	TitleId         string   `json:"titleId"`
	Ordering        int      `json:"ordering"`
	Title           string   `json:"title"`
	Region          string   `json:"region"`
	Language        string   `json:"language"`
	Types           []string `json:"types"`
	Attributes      []string `json:"attributes"`
	IsOriginalTitle bool     `json:"isOriginalTitle"`
}
