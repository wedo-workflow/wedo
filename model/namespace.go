package model

type Namespace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NamespaceQueryOptions is the query options for a namespace
type NamespaceQueryOptions struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NameLike string `json:"name_like"`
	// Limit is the maximum number of namespaces to return
	Limit int `json:"limit"`
	// Offset is the offset from the first namespace to return
	Offset int `json:"offset"`
}
