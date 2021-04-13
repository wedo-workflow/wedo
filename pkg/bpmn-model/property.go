package bpmn_model

type Property interface {
	SetVal(string, string) error
	Val(string) string
}

type Note struct {
}

func (n Note) SetVal(key string, val string) error {
	return nil
}

func (n Note) Val(key string) string {
	return ""
}

type ExtProperty struct {
}

func (e ExtProperty) SetVal(key string, val string) error {
	return nil
}

func (e ExtProperty) Val(key string) string {
	return ""
}
