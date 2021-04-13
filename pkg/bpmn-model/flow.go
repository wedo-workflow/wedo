package bpmn_model

type Flow interface {
	ToString() string
}

type Fork interface {
	Flow
}

type Join interface {
	Flow
}

type Decision interface {
	ToString() string
}

type ExclusiveDecision interface {
	Decision
}

type EventBasedDecision interface {
	Decision
}
type InclusiveDecision interface {
	Decision
}
type MergingDecision interface {
	Decision
}

type Looping interface {
	ToString() string
}

type ActivityLooping interface {
	Looping
}

type SequenceFlowLooping interface {
	Looping
}

type MultipleInstances interface {
}

type ProcessBreak interface {
}

type Transaction interface {
}

type OffPageConnector interface {
}
