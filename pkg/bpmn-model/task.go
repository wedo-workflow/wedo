package bpmn_model

type Task interface {
	Node
}

type TaskAtomic interface {
	Task
}

type TaskChoreography interface {
	Task
}
