package bpmn_model

// Connecting Sequence Flow
type Connecting interface {
	ToString() string
}

type SequenceFlow interface {
	Connecting
}

type NormalFlow interface {
	SequenceFlow
}

type UncontrolledFlow interface {
	SequenceFlow
}

type ConditionalFlow interface {
	SequenceFlow
}

type DefaultFlow interface {
	SequenceFlow
}

type ExceptionFlow interface {
	SequenceFlow
}

type CompensationAssociation interface {
	SequenceFlow
}

type MessageFlow interface {
	Connecting
}

type Association interface {
	Connecting
}

type DataAssociation interface {
	Connecting
}
