package element

type Swimlane interface {
	Element
}

type Pool interface {
	Swimlane
}

type Lane interface {
	Swimlane
}
