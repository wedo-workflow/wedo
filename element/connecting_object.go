package element

type ConnectingObject interface {
	Element
}

type SequenceFlow interface {
	ConnectingObject
}

type MessageFlow interface {
	ConnectingObject
}

type Association interface {
	ConnectingObject
}

type DataAssociation interface {
	ConnectingObject
}
