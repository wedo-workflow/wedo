package element

type Data interface {
	Element
}

type DataObject interface {
	Data
}

type DataInput interface {
	Data
}

type DataOutput interface {
	Data
}

type DataStore interface {
	Data
}
