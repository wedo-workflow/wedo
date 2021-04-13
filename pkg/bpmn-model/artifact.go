package bpmn_model

type Artifact interface {
	ToString() string
}

type Group interface {
	Artifact
}

type TextAnnotation interface {
	Artifact
}
