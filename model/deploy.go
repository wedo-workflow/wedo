package model

type Deploy struct {
	DID     string `json:"did"`
	Name    string `json:"name"`
	Content []byte `json:"content"`
}
