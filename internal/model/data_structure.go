package model

type Tree struct {
	Title    string `json:"title"`
	Value    string `json:"value,omitempty"`
	Key      string `json:"key,omitempty"`
	Children []Tree `json:"children,omitempty"`
}
