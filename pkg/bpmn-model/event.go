package bpmn_model

// Event represent all events happens in system.
type Event interface {
	Flow
	Node
	Type() EventType
}

// StartEvent Flow Dimension
type StartEvent interface {
	Event
	Listeners() []Listener
	ExtProperty() []Property
	Note() Property

	Start() error
}

// IntermediateEvent Flow Dimension
type IntermediateEvent interface {
	Event
	Listeners() []Listener
	ExtProperty() []Property
	Note() Property

	Intermediate() error
}

// EndEvent Flow Dimension
type EndEvent interface {
	Event
	Listeners() []Listener
	ExtProperty() []Property
	Note() Property

	End() error
}

// EventType Type Dimension
type EventType interface {
	ToString() string
}
