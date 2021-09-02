package element

type Artifact interface {
	Element
}

type Group interface {
	Artifact
}

type TextAnnotation interface {
	Artifact
}
