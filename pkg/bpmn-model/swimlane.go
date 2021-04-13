package bpmn_model

type Swimlane interface {
	ToString() string
}

type Pool interface {
	Swimlane
}

type Lane interface {
	Swimlane
}
