package model

type Deploy struct {
	DID     string
	Name    string `json:"name"`
	Content []byte `json:"content"`
}
