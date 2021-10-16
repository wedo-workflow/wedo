package model

import "time"

type Deploy struct {
	DID        string
	Name       string    `json:"name"`
	Content    []byte    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
