package element

type FlowObject interface {
	Element
}

type Event interface {
	FlowObject
}

type Activity interface {
	FlowObject
}

type Gateway interface {
	FlowObject
}
